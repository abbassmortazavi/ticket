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
		Name:      name,
		Clients:   make(map[*Client]bool),
		Join:      make(chan *Client),
		Leave:     make(chan *Client),
		Forward:   make(chan []byte),
		usernames: make(map[string]bool), // Initialize usernames map
	}
}

// Run starts the room's event loop
func (r *Room) Run() {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("Recovered from panic in room %s: %v", r.Name, rec)
		}
	}()

	for {
		select {
		case client := <-r.Join:
			if r.closed {
				continue
			}
			r.mu.Lock()
			r.Clients[client] = true
			r.usernames[client.Name] = true // Add to usernames
			r.mu.Unlock()

			// Send user join notification
			r.broadcastUserJoin(client.Name)
			// Send updated user list to all clients
			r.broadcastUserList()

		case client := <-r.Leave:
			r.mu.Lock()
			if _, exists := r.Clients[client]; exists {
				delete(r.Clients, client)
				delete(r.usernames, client.Name) // Remove from usernames
				select {
				case _, ok := <-client.Receive:
					if ok {
						close(client.Receive)
					}
				default:
					close(client.Receive)
				}
			}
			r.mu.Unlock()
			if !r.closed {
				// Send user leave notification
				r.broadcastUserLeave(client.Name)
				// Send updated user list to all clients
				r.broadcastUserList()
			}

		case message, ok := <-r.Forward:
			if !ok {
				return
			}
			r.mu.RLock()
			for client := range r.Clients {
				select {
				case client.Receive <- message:
				default:
					go func(c *Client) {
						r.Leave <- c
					}(client)
				}
			}
			r.mu.RUnlock()
		}
	}
}

// ServeHTTP handles WebSocket connections
func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.closed {
		http.Error(w, "Room is closed", http.StatusGone)
		return
	}

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("websocket upgrade error:", err)
		return
	}

	client := &Client{
		Socket:  conn,
		Receive: make(chan []byte, 256),
		Room:    r,
		Name:    generateUserName(),
	}

	// Send user_info to THIS client only
	client.Socket.WriteJSON(struct {
		Type     string `json:"type"`
		Username string `json:"username"`
	}{
		Type:     "user_info",
		Username: client.Name,
	})

	// Register client
	r.Join <- client

	// Send initial user list to this client
	r.sendUserList(client)

	// Start goroutines
	go client.write()
	go client.read()

	log.Printf("WebSocket connection established for room: %s, user: %s", r.Name, client.Name)
}

// Broadcast user list to all clients
func (r *Room) broadcastUserList() {
	userList := r.getUserList()
	message := Message{
		Type:  "user_list",
		Users: userList,
	}

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling user list: %v", err)
		return
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	for client := range r.Clients {
		select {
		case client.Receive <- jsonMsg:
		default:
			// Skip if channel is blocked
		}
	}
}

// Broadcast user join notification
func (r *Room) broadcastUserJoin(username string) {
	message := Message{
		Type:     "user_join",
		Username: username,
	}

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling user join: %v", err)
		return
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	for client := range r.Clients {
		select {
		case client.Receive <- jsonMsg:
		default:
			// Skip if channel is blocked
		}
	}
}

// Broadcast user leave notification
func (r *Room) broadcastUserLeave(username string) {
	message := Message{
		Type:     "user_leave",
		Username: username,
	}

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling user leave: %v", err)
		return
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	for client := range r.Clients {
		select {
		case client.Receive <- jsonMsg:
		default:
			// Skip if channel is blocked
		}
	}
}

// ClientCount returns the number of clients in the room
func (r *Room) ClientCount() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.Clients)
}

// Close closes the room and all connections
func (r *Room) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed {
		return // Already closed
	}

	r.closed = true

	// Close channels safely
	close(r.Join)
	close(r.Leave)
	close(r.Forward)

	// Close all client connections
	for client := range r.Clients {
		client.Socket.Close()
		select {
		case _, ok := <-client.Receive:
			if ok {
				close(client.Receive)
			}
		default:
			close(client.Receive)
		}
	}
	r.Clients = make(map[*Client]bool)
}

// broadcastSystemMessage sends system messages to all clients
func (r *Room) broadcastSystemMessage(msg string) {
	if r.closed {
		return
	}

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

	r.mu.RLock()
	defer r.mu.RUnlock()
	for client := range r.Clients {
		select {
		case client.Receive <- jsonMsg:
		default:
			// Skip if channel is blocked
		}
	}
}

// Get current list of usernames
func (r *Room) getUserList() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	usernames := make([]string, 0, len(r.usernames))
	for username := range r.usernames {
		usernames = append(usernames, username)
	}
	return usernames
}

// Send user list to a specific client
func (r *Room) sendUserList(client *Client) {
	userList := r.getUserList()
	message := Message{
		Type:  "user_list",
		Users: userList,
	}

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling user list: %v", err)
		return
	}

	select {
	case client.Receive <- jsonMsg:
	default:
		// Channel is full, skip
	}
}

// generateUserName generates a random user name
func generateUserName() string {
	return "user_" + string(rune(rand.Intn(1000)+65))
}
