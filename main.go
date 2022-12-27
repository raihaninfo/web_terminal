package main

import (
	"log"
	"net/http"

	"github.com/raihaninfo/web_terminal/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.Ws)
	http.HandleFunc("/", handlers.Home)

	// static files server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
