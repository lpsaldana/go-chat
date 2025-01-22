package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lpsaldana/go-chat/internal/models"
	"github.com/lpsaldana/go-chat/internal/services"
)

type ChatHandler struct {
	chatService services.ChatService
}

func NewChatHandler(chatService services.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (r *ChatHandler) FindAll(c *gin.Context) {
	chats, err := r.chatService.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "error",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"chats": chats,
	})
}

func (r *ChatHandler) SaveChat(c *gin.Context) {
	var newChat models.Chat
	if err := c.BindJSON(&newChat); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	chatID, err := r.chatService.Save(newChat)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "error",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"chatID": chatID,
	})
}
