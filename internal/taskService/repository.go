package taskService

import "gorm.io/gorm"

// MessageRepository - интерфейс для работы с репозиторием
type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id uint, message Message) (Message, error)
	DeleteMessageByID(id uint) error
}

// messageRepository - структура, которая реализует интерфейс MessageRepository
type messageRepository struct {
	db *gorm.DB
}

// NewMessageRepository - конструктор для создания репозитория
func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

// CreateMessage - создание сообщения в БД
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

// GetAllMessages - получение всех сообщений
func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

// UpdateMessageByID - обновление сообщения по ID
func (r *messageRepository) UpdateMessageByID(id uint, message Message) (Message, error) {
	result := r.db.Model(&Message{}).Where("id = ?", id).Updates(message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

// DeleteMessageByID - удаление сообщения по ID
func (r *messageRepository) DeleteMessageByID(id uint) error {
	result := r.db.Delete(&Message{}, id)
	return result.Error
}
