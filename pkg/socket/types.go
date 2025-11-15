package socket

import "github.com/gorilla/websocket"

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Room    string `json:"room,omitempty"`
	Type    string `json:"type,omitempty"` //message,join,leave,etc
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
}
