package taskService

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task   string `json:"task"`    // Наш сервер будет ожидать json с полем text
	IsDone bool   `json:"is_done"` // В Go используем CamelCase, в JSON - snake_case
}
