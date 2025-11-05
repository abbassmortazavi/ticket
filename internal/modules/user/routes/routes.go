package routes

import (
	userCtrl "ticket/internal/modules/user/controllers"

	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	userController := userCtrl.New()
	router.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/register", userController.CreateUser)
	})

}
