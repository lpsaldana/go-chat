package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/lpsaldana/go-chat/internal/models"
)

func TestHandleConnections_MissingRoomParameter(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ws", nil)
	rr := httptest.NewRecorder()

	HandleConnections(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestHandleConnections_ValidConnection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandleConnections))
	defer server.Close()

	u := "ws" + server.URL[4:] + "?room=testRoom"
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer ws.Close()
}

func TestBroadcastMessage(t *testing.T) {
	roomId := "testRoom"
	user := &models.User{Send: make(chan []byte, 1)}
	chatRooms[roomId] = map[*models.User]bool{user: true}

	message := []byte("Hello, Chat!")

	broadcastMessage(roomId, message)

	select {
	case msg := <-user.Send:
		if string(msg) != "Hello, Chat!" {
			t.Errorf("Expected %s, but got %s", "Hello, Chat!", msg)
		}
	default:
		t.Error("Message was not sent")
	}
}
