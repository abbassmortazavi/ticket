package controllers

import (
	"log"
	"net/http"
	"ticket/internal/modules/message/repositories/message"
	"ticket/internal/utils"
	"ticket/pkg/html"
	"ticket/pkg/session"
)

type Controller struct {
	messageRepository message.MessageRepositoryInterface
}

func New() *Controller {
	return &Controller{
		messageRepository: message.New(),
	}
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	userID, err := session.Get(r, "user_id")
	if err != nil {
		log.Println("session error:", err)
	}

	name, err := session.Get(r, "name")
	if err != nil {
		log.Println("session error:", err)
	}

	data := map[string]interface{}{
		"title":  "Home Page",
		"userID": userID,
		"name":   name,
	}

	html.Render(w, "index", data)
}
func (c *Controller) Chat(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Chat Page",
	}
	html.Render(w, "chat", data)
}

type SaveMessageRequest struct {
	Message string `json:"message"`
	Room    string `json:"room"`
}

func (c *Controller) SaveMessage(w http.ResponseWriter, r *http.Request) {
	var req SaveMessageRequest
	err := utils.ReadJson(w, r, &req)
	if err != nil {
		log.Println(err)
		return
	}

	userID, err := session.Get(r, "user_id")
	if err != nil {
		log.Println("session error:", err)
		return
	}
	name, err := session.Get(r, "name")
	if err != nil {
		log.Println("session error:", err)
		return
	}
	log.Println("name:", name)
	log.Println("userID:", userID)

}
