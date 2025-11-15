package routes

import (
	"net/http"
	auth "ticket/internal/modules/auth/routes"
	home "ticket/internal/modules/home/routes"
	ticket "ticket/internal/modules/ticket/routes"
	user "ticket/internal/modules/user/routes"
	"ticket/pkg/socket"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router) {
	ticket.Routes(router)
	user.Routes(router)
	auth.Routes(router)
	home.Routes(router)
	// Register WebSocket route directly
	router.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		roomName := r.URL.Query().Get("room")
		if roomName == "" {
			http.Error(w, "Room name is required", http.StatusBadRequest)
			return
		}

		room := socket.GetOrCreateRoom(roomName)
		room.ServeHTTP(w, r)
	})
}
