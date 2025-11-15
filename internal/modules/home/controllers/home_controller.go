package controllers

import (
	"net/http"
	"ticket/pkg/html"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Home Page",
	}
	html.Render(w, "index", data)
}
func (c *Controller) Chat(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Chat Page",
	}
	html.Render(w, "chat", data)
}
