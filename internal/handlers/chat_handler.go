package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lpsaldana/go-chat/internal/models"
)

var (
	chatRooms = make(map[string]map[*models.User]bool)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	mutex sync.Mutex
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	roomId := r.URL.Query().Get("room")
	if roomId == "" {
		http.Error(w, "Missing room parameter", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error updating websocket", err)
		return
	}
	defer conn.Close()

	client := &models.User{Conn: conn, Send: make(chan []byte)}
	go client.WritePump()
	mutex.Lock()
	if _, exists := chatRooms[roomId]; !exists {
		chatRooms[roomId] = make(map[*models.User]bool)
	}
	chatRooms[roomId][client] = true
	mutex.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message", err)
			return
		}
		fmt.Printf("Mensaje for chatroom %s: %s\n", roomId, msg)
		broadcastMessage(roomId, msg)
	}
}

func broadcastMessage(roomId string, message []byte) {
	mutex.Lock()
	defer mutex.Unlock()
	for user := range chatRooms[roomId] {
		select {
		case user.Send <- message:
			fmt.Println("Mensaje enviado al canal")
		default:
			fmt.Println("Cliente desconectado o canal lleno")
			close(user.Send)
			delete(chatRooms[roomId], user)
		}
	}
}
