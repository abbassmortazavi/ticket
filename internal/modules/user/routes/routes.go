package routes

import (
	userCtrl "ticket/internal/modules/user/controllers"
	"ticket/pkg/middlewares"

	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	userController := userCtrl.New()
	authMiddleware := middlewares.GetMiddleware()
	router.Route("/api/v1/users", func(r chi.Router) {
		r.Use(authMiddleware.AuthMiddleware)
		r.Get("/get-user/{id}", userController.GetUser)
	})

}
