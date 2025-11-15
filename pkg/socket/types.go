package socket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Name     string   `json:"name"`
	Message  string   `json:"message"`
	Room     string   `json:"room,omitempty"`
	Type     string   `json:"type,omitempty"`     // message, system, user_join, user_leave, user_list
	Username string   `json:"username,omitempty"` // For user_join/user_leave events
	Users    []string `json:"users,omitempty"`    // For user_list events
}

// Client represents a WebSocket client
type Client struct {
	Socket  *websocket.Conn
	Receive chan []byte
	Room    *Room
	Name    string
}

// Room represents a chat room
type Room struct {
	Name    string
	Clients map[*Client]bool
	Join    chan *Client
	Leave   chan *Client
	Forward chan []byte
	mu      sync.RWMutex // Add mutex for thread safety
	closed  bool         // Track if room is closed
	// Add user tracking
	usernames map[string]bool // Track usernames for the user list
}
