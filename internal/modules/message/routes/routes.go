package routes

import (
	mesaageCtl "ticket/internal/modules/message/controllers"

	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	messageController := mesaageCtl.New()
	router.Post("/message", messageController.SaveMessage)

}
