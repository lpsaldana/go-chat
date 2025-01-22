package services

import (
	"github.com/lpsaldana/go-chat/internal/models"
	"github.com/lpsaldana/go-chat/internal/repositories"
)

type ChatService interface {
	FindAll() (*[]models.Chat, error)
	Save(chat models.Chat) (string, error)
}

type chatService struct {
	chatRepository repositories.ChatRepository
}

func NewChatService(chatRepository repositories.ChatRepository) ChatService {
	return &chatService{chatRepository: chatRepository}
}

func (r *chatService) FindAll() (*[]models.Chat, error) {
	return r.chatRepository.FindAll()
}

func (r *chatService) Save(chat models.Chat) (string, error) {
	return r.chatRepository.SaveChat(chat)
}
