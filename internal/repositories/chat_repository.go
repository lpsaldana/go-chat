package repositories

import "github.com/lpsaldana/go-chat/internal/models"

type ChatRepository interface {
	FindAll() (*[]models.Chat, error)
	SaveChat(chat models.Chat) (string, error)
}

type chatRepository struct {
	chat []models.Chat
}

func NewChatRepository() ChatRepository {
	return &chatRepository{chat: make([]models.Chat, 0)}
}

func (r *chatRepository) FindAll() (*[]models.Chat, error) {
	return &r.chat, nil
}

func (r *chatRepository) SaveChat(chat models.Chat) (string, error) {
	r.chat = append(r.chat, chat)
	return chat.ID, nil
}
