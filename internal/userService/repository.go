package userService

import (
	"fmt"
	"gorm.io/gorm"
)

// UserRepository - интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
	GetUserByEmail(email string) (User, error)
}

// userRepository - структура, которая реализует интерфейс UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository - конструктор для создания репозитория
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser - создание нового пользователя
func (r *userRepository) CreateUser(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// GetAllUsers - получение всех пользователей
func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID - получение пользователя по ID
func (r *userRepository) GetUserByID(id uint) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// UpdateUserByID - обновление пользователя по ID
func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User

	if err := r.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		if existingUser.ID != id {
			return User{}, fmt.Errorf("user with email %s already exists", user.Email)
		}
	} else if err != gorm.ErrRecordNotFound {
		return User{}, err
	}
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}

	if err := r.db.Model(&existingUser).Updates(user).Error; err != nil {
		return User{}, err
	}
	return existingUser, nil
}

// DeleteUserByID - удаление пользователя по ID
func (r *userRepository) DeleteUserByID(id uint) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *userRepository) GetUserByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
