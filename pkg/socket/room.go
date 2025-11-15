package socket

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  viper.GetInt("SOCKET_BUFFER_SIZE"),
		WriteBufferSize: viper.GetInt("MESSAGE_BUFFER_SIZE"),
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	roomsLock sync.Mutex
	rooms     = make(map[string]*Room)
)

func NewRoom(name string) *Room {
	return &Room{
		Name:    name,
		Clients: make(map[*Client]bool),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Forward: make(chan []byte),
	}
}

// Run starts the room's event loop
func (r *Room) Run() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
			r.broadcastSystemMessage(client.Name + " joined the room")
		case client := <-r.Leave:
			delete(r.Clients, client)
			close(client.Receive)
			r.broadcastSystemMessage(client.Name + " left the room")
		case message := <-r.Forward:
			for client := range r.Clients {
				select {
				case client.Receive <- message:
				default:
					delete(r.Clients, client)
					close(client.Receive)
				}
			}
		}

	}
}

// ServeHTTP handles WebSocket connections
func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("websocket upgrade error:", err)
		return
	}
	client := &Client{
		Socket:  conn,
		Receive: make(chan []byte),
		Room:    r,
		Name:    generateUserName(),
	}
	r.Join <- client
	defer func() {
		r.Leave <- client
	}()
	go client.write()
	client.read()
}

// ClientCount returns the number of clients in the room
func (r *Room) ClientCount() int {
	return len(r.Clients)

}

// Close closes the room and all connections
func (r *Room) Close() {
	for client := range r.Clients {
		client.Socket.Close()
		close(client.Receive)
	}
	r.Clients = make(map[*Client]bool)
}

// broadcastSystemMessage sends system messages to all clients
func (r *Room) broadcastSystemMessage(msg string) {
	systemMessage := Message{
		Name:    "System",
		Message: msg,
		Type:    "system",
	}
	jsonMsg, err := json.Marshal(systemMessage)
	if err != nil {
		log.Println("Send System Message Failed:", err)
		return
	}

	for client := range r.Clients {
		select {
		case client.Receive <- jsonMsg:
		default:
			delete(r.Clients, client)
			close(client.Receive)
		}
	}
}

// generateUserName generates a random user name
func generateUserName() string {
	return "user_" + string(rune(rand.Intn(1000)+65))
}
