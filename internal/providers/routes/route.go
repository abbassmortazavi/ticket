package routes

import (
	ticket "ticket/internal/modules/ticket/routes"
	user "ticket/internal/modules/user/routes"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router) {
	ticket.Routes(router)
	user.Routes(router)
}
