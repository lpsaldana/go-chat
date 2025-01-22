package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lpsaldana/go-chat/internal/bootstrap"
)

func main() {
	dependencies, _ := bootstrap.InitializeApp()
	fmt.Println("test chat")
	router := gin.Default()
	router.GET("/", dependencies.ChatHandler.FindAll)
	router.GET("/users", dependencies.UserHandler.FindAll)
	router.POST("/new-chat", dependencies.ChatHandler.SaveChat)
	router.Run(":8080")
}
