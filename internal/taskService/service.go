package taskService

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// MessageService - структура для работы с сообщениями
type MessageService struct {
	repo MessageRepository
}

// NewService - конструктор для создания нового сервиса
func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// CreateMessage - создание сообщения через сервис
func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

// GetAllMessages - получение всех сообщений через сервис
func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

var ErrMessageNotFound = errors.New("message not found")

func (s *MessageService) GetMessageByID(id uint) (Message, error) {
	message, err := s.repo.GetMessageByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Message{}, fmt.Errorf("message with ID %d not found", id)
		}
		return Message{}, fmt.Errorf("error retrieving message: %w", err)
	}
	return message, nil
}

// UpdateMessageByID - обновление сообщения через сервис
func (s *MessageService) UpdateMessageByID(id uint, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, message)
}

// DeleteMessageByID - удаление сообщения через сервис
func (s *MessageService) DeleteMessageByID(id uint) error {

	return s.repo.DeleteMessageByID(id)
}

// GetMessagesByUserID - получение всех задач для конкретного пользователя через сервис
func (s *MessageService) GetMessagesByUserID(userID uint) ([]Message, error) {
	return s.repo.GetMessagesByUserID(userID)
}
