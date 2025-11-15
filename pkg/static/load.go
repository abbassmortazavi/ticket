package static

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func LoadAsset(router *chi.Mux) {
	// Get absolute path to assets directory
	assetsPath := filepath.Join(".", "assets")

	// Create file server with proper path handling
	fs := http.FileServer(http.Dir(assetsPath))

	// Use StripPrefix to remove /assets from the request path
	router.Handle("/assets/*", http.StripPrefix("/assets", fs))
}
