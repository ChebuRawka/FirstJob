package handlers

import (
	"FirstJobProject/internal/userService" // Убедитесь, что это правильный путь
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// UserHandler - структура для обработки пользователей
type UserHandler struct {
	Service *userService.UserService
}

// NewUserHandler - создаёт новый обработчик для пользователей
func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// GetUsers - возвращает всех пользователей
func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving users"})
	}
	return c.JSON(http.StatusOK, users)
}

// PostUser - создает нового пользователя
func (h *UserHandler) PostUser(c echo.Context) error {
	var user userService.User

	// Биндим данные из запроса
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Проверка на наличие email и пароля
	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email and password are required"})
	}

	// Создаем пользователя
	createdUser, err := h.Service.CreateUser(user)
	if err != nil {
		log.Println("Error creating user:", err)
		if strings.Contains(err.Error(), "already exists") {
			return c.JSON(http.StatusConflict, map[string]string{"message": "User with this email already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating user"})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

// PatchUserByID - обновляет пользователя по ID
func (h *UserHandler) PatchUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}
	var user userService.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	updatedUser, err := h.Service.UpdateUserByID(uint(id), user)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return c.JSON(http.StatusConflict, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error updating user"})
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUserByID - удаляет пользователя по ID
func (h *UserHandler) DeleteUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}
	err = h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting user"})
	}
	return c.NoContent(http.StatusNoContent)
}
