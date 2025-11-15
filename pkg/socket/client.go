package socket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// read reads messages from the WebSocket connection
func (c *Client) read() {
	defer func() {
		if c.Room != nil {
			c.Room.Leave <- c
		}
		c.Socket.Close()
	}()

	for {
		_, rawMsg, err := c.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}

		// پیام خام کلاینت را پارس می‌کنیم
		var incoming struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		}

		if err := json.Unmarshal(rawMsg, &incoming); err != nil {
			log.Println("Invalid incoming JSON:", err)
			continue
		}

		// ساختن پیام نهایی برای همه
		msg := Message{
			Name:    c.Name,
			Message: incoming.Message, // فقط متن پیام
			Room:    c.Room.Name,
			Type:    "message",
		}

		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("JSON marshal error: %v", err)
			continue
		}

		if c.Room != nil {
			c.Room.Forward <- jsonMsg
		}
	}
}

// write writes messages to the WebSocket connection
func (c *Client) write() {
	defer c.Socket.Close()
	for {
		select {
		case msg, ok := <-c.Receive:
			if !ok {
				// Channel closed, send close message
				err := c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					return
				}
				return
			}
			err := c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				return
			}
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
