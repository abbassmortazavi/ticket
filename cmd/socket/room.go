package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

type Room struct {

	//hold all current clients in the Room
	clients map[*Client]bool

	//join is channel for all clients wishing to join this Room
	join chan *Client

	//leave the channel for all clients wishing to join this Room/
	leave chan *Client

	//forward is that channel holds incomming message that should be other clients
	forward chan []byte
}

func NewRoom() *Room {
	return &Room{
		clients: make(map[*Client]bool),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		forward: make(chan []byte),
	}
}
func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.receive)
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}

		}

	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: messageBufferSize,
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	roomName := req.URL.Query().Get("room")
	if roomName == "" {
		//roomName = "default"
		log.Println("room name is empty")
		http.Error(w, "room name is empty", http.StatusBadRequest)
	}
	realRoomName := GetRoom(roomName)
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("server http", err)
		return
	}
	client := &Client{
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    r,
		name:    fmt.Sprintf("user_%d", rand.Intn(1000)),
	}
	realRoomName.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

var rooms = make(map[string]*Room)
var roomsLock sync.Mutex

func GetRoom(name string) *Room {

	//prevent creating the same room name when multiple users do that the same time
	roomsLock.Lock()
	defer roomsLock.Unlock()
	//if the room name already exist
	if room, ok := rooms[name]; ok {
		return room
	}
	//if not
	room := NewRoom()
	rooms[name] = room
	go room.Run()
	return room
}
