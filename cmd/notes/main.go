package main

import (
	"log"
	"net/http"
	"notes/embedded"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	ADDR = "localhost:8080"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ws", wsHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.FS(embedded.FS)))

	srv := &http.Server{
		Handler:      r,
		Addr:         ADDR,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on", ADDR)
	log.Fatal(srv.ListenAndServe())
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	if err != nil {
		log.Println(err)
		return
	}

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%s sent: %s", ws.RemoteAddr(), message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println(err)
			return
		}
	}

}
