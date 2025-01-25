package userService

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// UserService - структура, которая реализует бизнес-логику для пользователей
type UserService struct {
	userRepo UserRepository
}

// NewUserService - конструктор для создания сервиса
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser - создание нового пользователя
func (s *UserService) CreateUser(user User) (User, error) {
	// Проверяем, существует ли пользователь с таким email
	existingUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, err
	}
	if existingUser.ID != 0 {
		return User{}, fmt.Errorf("user with email %s already exists", user.Email)
	}

	// Если пользователя нет, создаем нового
	return s.userRepo.CreateUser(user)
}

// GetAllUsers - получение всех пользователей
func (s *UserService) GetAllUsers() ([]User, error) {
	return s.userRepo.GetAllUsers()
}

// GetUserByID - получение пользователя по ID
func (s *UserService) GetUserByID(id uint) (User, error) {
	return s.userRepo.GetUserByID(id)
}

// UpdateUserByID - обновление пользователя по ID
func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.userRepo.UpdateUserByID(id, user)
}

// DeleteUserByID - удаление пользователя по ID
func (s *UserService) DeleteUserByID(id uint) error {
	return s.userRepo.DeleteUserByID(id)
}
