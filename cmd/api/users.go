package api

import (
	"encoding/json"
	"net/http"
)

func (app *Application) GetUser(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
	// Use w.Write() to send response
	// Use r to access the request
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Sample user data
	user := map[string]interface{}{
		"id":    1,
		"name":  "John Doe",
		"email": "john@example.com",
	}
	app.Logger.Info().Msgf("im here")

	// Encode to JSON and send response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
