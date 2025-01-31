package taskService

import (
	"FirstJobProject/internal/userService"
	"fmt"
	"gorm.io/gorm"
)

// MessageRepository - интерфейс для работы с репозиторием
type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	GetMessageByID(id uint) (Message, error) // Добавляем метод для получения сообщения по ID
	UpdateMessageByID(id uint, message Message) (Message, error)
	DeleteMessageByID(id uint) error
	GetMessagesByUserID(id uint) ([]Message, error)
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
	// Проверяем, существует ли пользователь с таким ID
	var user userService.User
	if err := r.db.First(&user, message.UserID).Error; err != nil {
		// Если пользователь не найден, возвращаем ошибку
		return Message{}, fmt.Errorf("user with ID %d not found", message.UserID)
	}

	// Если ID уже существует в запросе, то используем его. В противном случае, GORM сам установит ID.
	if message.ID == 0 {
		// Если ID не передан (или равен 0), то GORM будет использовать автоинкремент
		if err := r.db.Create(&message).Error; err != nil {
			return Message{}, err
		}
	} else {
		// Если ID передан вручную, то проверим, не существует ли уже записи с таким ID
		var existingMessage Message
		if err := r.db.First(&existingMessage, message.ID).Error; err == nil {
			// Если запись с таким ID уже существует, возвращаем ошибку
			return Message{}, fmt.Errorf("message with ID %d already exists", message.ID)
		}
		// Вставляем сообщение с вручную указанным ID
		if err := r.db.Create(&message).Error; err != nil {
			return Message{}, err
		}
	}

	return message, nil
}

// GetAllMessages - получение всех сообщений
// GetAllMessages - получение всех сообщений с подгрузкой пользователя
func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Where("deleted_at IS NULL").Find(&messages).Error
	return messages, err
}

// GetMessageByID - получение сообщения по ID
func (r *messageRepository) GetMessageByID(id uint) (Message, error) {
	var message Message
	if err := r.db.First(&message, id).Error; err != nil {
		return Message{}, fmt.Errorf("message with ID %d not found", id)
	}
	return message, nil
}

// UpdateMessageByID - обновление сообщения по ID
// UpdateMessageByID - обновление сообщения по ID
func (r *messageRepository) UpdateMessageByID(id uint, message Message) (Message, error) {
	// Проверяем, существует ли сообщение с таким ID
	var existingMessage Message
	if err := r.db.First(&existingMessage, id).Error; err != nil {
		// Если запись не найдена, возвращаем ошибку
		return Message{}, fmt.Errorf("message with ID %d not found", id)
	}

	// Проверяем, существует ли пользователь с таким user_id
	var user userService.User
	if err := r.db.First(&user, message.UserID).Error; err != nil {
		return Message{}, fmt.Errorf("user with ID %d not found", message.UserID)
	}

	// Обновляем только те поля, которые переданы
	updates := map[string]interface{}{
		"task":    message.Task,
		"is_done": message.IsDone,
	}

	// Если поле "user_id" обновляется, добавляем его в обновления
	if message.UserID != 0 {
		updates["user_id"] = message.UserID
	}

	// Обновляем запись в базе данных
	if err := r.db.Model(&existingMessage).Updates(updates).Error; err != nil {
		return Message{}, err
	}

	// Возвращаем обновленную задачу
	return existingMessage, nil
}

// DeleteMessageByID - удаление сообщения по ID
func (r *messageRepository) DeleteMessageByID(id uint) error {
	// Проверяем, существует ли сообщение с таким ID
	var existingMessage Message
	if err := r.db.First(&existingMessage, id).Error; err != nil {
		// Если запись не найдена, возвращаем ошибку
		return fmt.Errorf("message with ID %d not found", id)
	}

	// Если запись существует, удаляем её
	if err := r.db.Delete(&Message{}, id).Error; err != nil {
		return err
	}

	return nil
}

// GetMessagesByUserID - получение всех задач для конкретного пользователя
func (r *messageRepository) GetMessagesByUserID(userID uint) ([]Message, error) {
	var messages []Message
	err := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&messages).Error
	return messages, err
}
