package routes

import (
	"github.com/go-chi/chi/v5"
	authCtrl "ticket/internal/modules/auth/controllers"
)

func Routes(router chi.Router) {
	authController := authCtrl.New()
	router.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", authController.Login)
		r.Post("/register", authController.Register)
	})

}
