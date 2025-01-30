package userService

import (
	"gorm.io/gorm"
)

// Реализация интерфейса UserInterface
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}
