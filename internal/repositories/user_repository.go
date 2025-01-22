package repositories

import (
	"github.com/lpsaldana/go-chat/internal/models"
)

type UserRepository interface {
	FindAll() (*[]models.User, error)
}

type userRepository struct {
	users []models.User
}

func NewUserRepository() UserRepository {
	var firstUsers = []models.User{
		{ID: "001", Username: "Leo"},
		{ID: "002", Username: "Ana"},
	}
	return &userRepository{users: firstUsers}
}

func (r *userRepository) FindAll() (*[]models.User, error) {
	return &r.users, nil
}
