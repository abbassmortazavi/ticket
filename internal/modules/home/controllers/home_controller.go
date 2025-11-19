package controllers

import (
	"context"
	"log"
	"net/http"
	"ticket/internal/modules/message/models"
	"ticket/internal/modules/message/repositories/message"
	"ticket/internal/modules/message/repositories/room"
	"ticket/internal/utils"
	"ticket/pkg/html"
	"ticket/pkg/session"
)

type Controller struct {
	messageRepository message.MessageRepositoryInterface
	roomRepository    room.RoomRepositoryInterface
}

func New() *Controller {
	return &Controller{
		messageRepository: message.New(),
		roomRepository:    room.New(),
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
	ctx := context.Background()
	roomFind, err := c.roomRepository.FindRoom(ctx, req.Room)
	if err != nil {
		log.Println(err)
		return
	}
	if roomFind == nil {
		log.Println("room not found")
		return
	}

	userID, err := session.Get(r, "user_id")
	if err != nil {
		log.Println("session error:", err)
		return
	}
	// Type assert to int
	user, ok := userID.(int)
	if !ok {
		log.Println("user_id is not an integer")
		return
	}

	messageSave := models.Message{
		RoomID:      roomFind.ID,
		UserID:      user,
		Message:     req.Message,
		MessageType: "text",
	}

	res, err := c.messageRepository.SaveMessage(ctx, &messageSave)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res)
}
