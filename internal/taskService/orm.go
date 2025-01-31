package taskService

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`

	//User   userService.User `json:"user"` // здесь используем конкретную структуру User
}

//type MessageWithUser struct {
//	Message
//	User userService.User `json:"user"` // Добавляем информацию о пользователе
//}

// TaskWithoutUser - используется для GET /users/:id/tasks
type TaskWithUserID struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	UserID    uint      `json:"user_id"` // Включаем user_id
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TaskWithoutUser - используется для GET /users/:id/tasks
type TaskWithoutUser struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
