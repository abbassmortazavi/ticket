package controllers

import (
	"log"
	"net/http"
	"ticket/internal/modules/message/Requests"
	"ticket/internal/modules/message/models"
	"ticket/internal/modules/message/repositories/message"
	"ticket/internal/utils"
)

type Controller struct {
	messageRepo message.MessageRepositoryInterface
}

func New() *Controller {
	return &Controller{
		messageRepo: message.New(),
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
