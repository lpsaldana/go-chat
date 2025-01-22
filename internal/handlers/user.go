package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lpsaldana/go-chat/internal/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (r *UserHandler) FindAll(c *gin.Context) {
	users, err := r.service.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error fetching users"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"users": users,
	})
}
