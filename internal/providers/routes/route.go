package routes

import (
	ticket "ticket/internal/modules/ticket/routes"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router) {
	ticket.Routes(router)
}
