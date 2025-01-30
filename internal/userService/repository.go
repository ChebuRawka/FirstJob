package userService

import (
	"fmt"
	"gorm.io/gorm"
	"log"
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
	// Проверяем, существует ли пользователь с таким email
	existingUser, err := r.GetUserByEmail(user.Email)
	if err == nil && existingUser.ID != 0 {
		log.Println("User with email already exists:", user.Email)
		return User{}, fmt.Errorf("user with email %s already exists", user.Email)
	}

	// Создаем нового пользователя с задачами
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, fmt.Errorf("error creating user: %v", err)
	}

	return user, nil
}

// GetAllUsers - получение всех пользователей
func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	// Указываем модель явно
	if err := r.db.Model(&User{}).Find(&users).Error; err != nil {
		log.Println("Error fetching users:", err)
		return nil, fmt.Errorf("error fetching users: %v", err)
	}

	return users, nil
}

// GetUserByID - получение пользователя по ID
func (r *userRepository) GetUserByID(id uint) (User, error) {
	var user User
	// Загружаем пользователя вместе с его задачами
	if err := r.db.Preload("Messages").First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// UpdateUserByID - обновление пользователя по ID
func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User

	// Указываем таблицу, с которой работаем
	if err := r.db.Model(&User{}).Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		if existingUser.ID != id {
			return User{}, fmt.Errorf("user with email %s already exists", user.Email)
		}
	} else if err != gorm.ErrRecordNotFound {
		return User{}, err
	}

	// Указываем таблицу для обновления
	if err := r.db.Model(&User{}).First(&existingUser, id).Error; err != nil {
		return User{}, err
	}

	if err := r.db.Model(&existingUser).Updates(user).Error; err != nil {
		return User{}, err
	}
	return existingUser, nil
}

// DeleteUserByID - удаление пользователя по ID
func (r *userRepository) DeleteUserByID(id uint) error {
	// Вместо того чтобы передавать &User{}, используй саму структуру User{}
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (User, error) {
	var user User
	// Здесь мы должны использовать .First для поиска одного пользователя по email
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return User{}, nil // Если пользователь не найден, возвращаем пустую структуру
		}
		return User{}, fmt.Errorf("error fetching user by email: %v", err)
	}
	return user, nil
}
