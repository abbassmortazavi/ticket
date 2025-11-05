package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	router.Route("/tickets", func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			response := map[string]interface{}{
				"success": true,
				"tickets": nil,
				"data":    "this is a test",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

		})
	})
}
