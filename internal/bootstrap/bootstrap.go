package bootstrap

import (
	"github.com/lpsaldana/go-chat/internal/handlers"
	"github.com/lpsaldana/go-chat/internal/repositories"
	"github.com/lpsaldana/go-chat/internal/services"
)

type AppDependencies struct {
	UserHandler *handlers.UserHandler
	ChatHandler *handlers.ChatHandler
}

func InitializeApp() (*AppDependencies, error) {
	//user
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHander := handlers.NewUserHandler(userService)

	//chat
	chatRepository := repositories.NewChatRepository()
	chatService := services.NewChatService(chatRepository)
	chatHandler := handlers.NewChatHandler(chatService)

	return &AppDependencies{UserHandler: userHander, ChatHandler: chatHandler}, nil
}
