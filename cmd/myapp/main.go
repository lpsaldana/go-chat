package main

import (
	"log"
	"net/http"

	"github.com/lpsaldana/go-chat/internal/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.HandleConnections)
	log.Println("WebSocket server port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
