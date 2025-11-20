package controllers

import (
	"context"
	"log"
	"net/http"
	"ticket/internal/modules/message/Requests"
	"ticket/internal/modules/message/models"
	"ticket/internal/modules/message/repositories/message"
	"ticket/internal/modules/message/repositories/room"
	response2 "ticket/internal/modules/message/response"
	"ticket/internal/modules/user/services"
	"ticket/internal/utils"
	"ticket/pkg/session"
)

type Controller struct {
	messageRepo    message.MessageRepositoryInterface
	userService    services.UserServiceInterface
	roomRepository room.RoomRepositoryInterface
}

func New() *Controller {
	return &Controller{
		messageRepo:    message.New(),
		userService:    services.New(),
		roomRepository: room.New(),
	}
}

func (c *Controller) SaveMessage(w http.ResponseWriter, r *http.Request) {
	var req Requests.MessageRequest
	err := utils.ReadJson(w, r, &req)
	if err != nil {
		log.Println("validation error:", err)
		utils.BadRequest(w, "validation error:", err)
		return
	}
	if !utils.ValidateStruct(w, &req) {
		return
	}
	ctx := r.Context()
	data := models.Message{
		UserID:      req.UserID,
		RoomID:      req.RoomID,
		MessageType: req.MessageType,
		Message:     req.Message,
	}
	res, err := c.messageRepo.SaveMessage(ctx, &data)
	if err != nil {
		log.Println("save message error:", err)
		utils.InternalError(w, err, "save message error.")
		return
	}
	log.Println("save message success:", res)
	utils.Success(w, http.StatusOK, res, "message saved")

}

func (c *Controller) GetMessages(w http.ResponseWriter, r *http.Request) {
	roomName := r.URL.Query().Get("room")
	if roomName == "" {
		log.Println("room not found")
		return
	}
	ctx := context.Background()
	roomFind, err := c.roomRepository.FindRoom(ctx, roomName)
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
	currentUserID, ok := userID.(int)
	if !ok {
		log.Println("user_id is not an integer")
		return
	}

	messages, err := c.messageRepo.GetMessagesByRoomId(ctx, roomFind.ID)
	if err != nil {
		log.Println(err)
		return
	}
	var response []response2.MessageResponse
	for _, msg := range messages {
		user, err := c.userService.GetById(ctx, msg.UserID)
		var username string
		if err != nil {
			log.Println(err)
			username = "unknown"
		} else {
			username = user.FullName
		}
		response = append(response, response2.MessageResponse{
			ID:          msg.ID,
			Username:    username,
			Message:     msg.Message,
			MessageType: msg.MessageType,
			IsOwn:       msg.ID == currentUserID,
			CreatedAt:   msg.CreatedAt,
		})
	}
	data := map[string]interface{}{
		"messages":        response,
		"success":         true,
		"current_user_id": currentUserID,
	}
	err = utils.WriteJson(w, http.StatusOK, data)
	if err != nil {
		return

	}
}
