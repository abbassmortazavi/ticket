package socket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	socket  *websocket.Conn
	receive chan []byte

	room *Room
	name string
}

// send message
func (c *Client) read() {
	defer c.socket.Close()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		ougoing := map[string]string{
			"name":    c.name,
			"message": string(message),
		}
		jsonOutgoing, err := json.Marshal(ougoing)
		if err != nil {
			log.Println("encoding faild", err)
			return
		}

		c.room.forward <- jsonOutgoing

	}
}
func (c *Client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
