package bootstrap

import (
	"github.com/lpsaldana/go-chat/internal/handlers"
	"github.com/lpsaldana/go-chat/internal/repositories"
	"github.com/lpsaldana/go-chat/internal/services"
)

type AppDependencies struct {
	UserHandler *handlers.UserHandler
}

func InitializeApp() (*AppDependencies, error) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHander := handlers.NewUserHandler(userService)

	return &AppDependencies{UserHandler: userHander}, nil
}
