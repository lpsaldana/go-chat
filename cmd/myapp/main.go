package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lpsaldana/go-chat/internal/bootstrap"
)

type chat struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  string `json:"userid"`
}

var chats = []chat{}

func getChats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"chats": chats,
	})
}

func postChat(c *gin.Context) {
	var newChat chat
	if err := c.BindJSON(&newChat); err != nil {
		return
	}
	chats = append(chats, newChat)
	c.IndentedJSON(http.StatusCreated, gin.H{
		"newChat": newChat,
	})
}

func main() {
	dependencies, _ := bootstrap.InitializeApp()
	fmt.Println("test chat")
	router := gin.Default()
	router.GET("/", getChats)
	router.GET("/users", dependencies.UserHandler.FindAll)
	router.POST("/new-chat", postChat)
	router.Run(":8080")
}
