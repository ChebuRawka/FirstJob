package userService

import "gorm.io/gorm"

// User - структура пользователя
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
