package services

import (
	"github.com/lpsaldana/go-chat/internal/models"
	"github.com/lpsaldana/go-chat/internal/repositories"
)

type UserService interface {
	FindAllUsers() (*[]models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (r *userService) FindAllUsers() (*[]models.User, error) {
	return r.userRepo.FindAll()
}
