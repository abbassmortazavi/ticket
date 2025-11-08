package routing

import (
	"ticket/internal/providers/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() {
	router = chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}

func GetRouter() *chi.Mux {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
