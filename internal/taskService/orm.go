package taskService

import (
	"FirstJobProject/internal/userService" // импортируем нужный пакет
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
	//User   userService.User `json:"user"` // здесь используем конкретную структуру User
}

type MessageWithUser struct {
	Message
	User userService.User `json:"user"` // Добавляем информацию о пользователе
}
