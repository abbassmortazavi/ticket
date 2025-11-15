package socket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// read reads messages from the WebSocket connection
func (c *Client) read() {
	defer c.Socket.Close()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				return
			}
		}
		msg := Message{
			Name:    c.Name,
			Message: string(message),
			Room:    c.Room.Name,
			Type:    "message",
		}

		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		c.Room.Forward <- jsonMsg
	}
}

// write writes messages to the WebSocket connection
func (c *Client) write() {
	defer c.Socket.Close()
	for msg := range c.Receive {
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
	}
}

// GetName returns the client's name
func (c *Client) GetName() string {
	return c.Name
}

// GetRoom returns the client's room
func (c *Client) GetRoom() *Room {
	return c.Room
}
