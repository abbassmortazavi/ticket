package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"ticket/pkg/socket"
)

type TemplateHandler struct {
	template *template.Template
	once     sync.Once
	filename string
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.template.Execute(w, r)
}

func main() {

	//rabbitmq.Send()
	//rabbitmq.Receive()
	var addr = flag.String("addr", "localhost:8083", "http service address")
	flag.Parse()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", &TemplateHandler{
		filename: "index.html",
	})
	http.Handle("/chat", &TemplateHandler{
		filename: "chat.html",
	})
	// WebSocket endpoint - using the global function for backward compatibility
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		roomName := r.URL.Query().Get("room")
		if roomName == "" {
			http.Error(w, "Room name is required", http.StatusBadRequest)
			return
		}

		room := socket.GetOrCreateRoom(roomName)
		room.ServeHTTP(w, r)
	})

	log.Printf("listening on %s", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Println("http service listen error:", err)
		panic(err)
	}
	//	cmd.Execute()
}
