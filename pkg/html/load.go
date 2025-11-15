package html

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

var templates *template.Template

func LoadHtml(router *chi.Mux) {
	var files []string

	// Walk through all directories to find HTML files
	err := filepath.Walk("internal", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			files = append(files, path)
			log.Printf("Found template: %s", path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	templates = template.Must(template.ParseFiles(files...))

	// Log all defined template names
	log.Printf("Loaded templates:")
	for _, t := range templates.Templates() {
		log.Printf("  - %s", t.Name())
	}
}

// RenderTemplate Helper function for handlers
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
