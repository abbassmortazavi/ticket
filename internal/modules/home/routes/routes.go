package routes

import (
	homeCtrl "ticket/internal/modules/home/controllers"

	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	homeController := homeCtrl.New()
	router.Get("/join", homeController.Home)
	router.Get("/chat", homeController.Chat)
}
