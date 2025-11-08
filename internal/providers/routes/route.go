package routes

import (
	auth "ticket/internal/modules/auth/routes"
	ticket "ticket/internal/modules/ticket/routes"
	user "ticket/internal/modules/user/routes"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router) {
	ticket.Routes(router)
	user.Routes(router)
	auth.Routes(router)
}
