package routes

import (
	authCtrl "ticket/internal/modules/auth/controllers"

	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	authController := authCtrl.New()
	router.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", authController.Login)
		r.Post("/register", authController.Register)
	})

}
