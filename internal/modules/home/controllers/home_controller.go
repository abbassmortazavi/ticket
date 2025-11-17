package controllers

import (
	"log"
	"net/http"
	"ticket/pkg/html"
	"ticket/pkg/session"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	userID, err := session.Get(r, "user_id")
	if err != nil {
		log.Println("session error:", err)
	}

	name, err := session.Get(r, "name")
	if err != nil {
		log.Println("session error:", err)
	}

	data := map[string]interface{}{
		"title":  "Home Page",
		"userID": userID,
		"name":   name,
	}

	html.Render(w, "index", data)
}
func (c *Controller) Chat(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Chat Page",
	}
	html.Render(w, "chat", data)
}
