package main

import (
	"fmt"
	"log"
	"net/http"
	"notes/embedded"
	"time"

	"github.com/gorilla/mux"
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

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")

}
