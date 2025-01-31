package handlers

import (
	"FirstJobProject/internal/taskService"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// TaskHandler - структура обработчика для задач
type TaskHandler struct {
	Service *taskService.MessageService
}

// NewTaskHandler - конструктор для создания нового обработчика
func NewTaskHandler(service *taskService.MessageService) *TaskHandler {
	return &TaskHandler{Service: service}
}

// GetTasks - возвращает все задачи
func (h *TaskHandler) GetTasks(c echo.Context) error {
	allTasks, err := h.Service.GetAllMessages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving tasks"})
	}

	// Возвращаем задачи с подгруженными пользователями
	return c.JSON(http.StatusOK, allTasks)
}

// GetTasksByUserID - возвращает все задачи для конкретного пользователя
func (h *TaskHandler) GetTasksByUserID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	tasks, err := h.Service.GetMessagesByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving tasks"})
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "No tasks found for this user"})
	}

	return c.JSON(http.StatusOK, tasks)
}

// PostTasks - создает новую задачу
func (h *TaskHandler) PostTasks(c echo.Context) error {
	var task taskService.Message
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Создаем задачу через сервис
	createdTask, err := h.Service.CreateMessage(task)
	if err != nil {
		// Если ошибка — проверяем, была ли ошибка о несуществующем пользователе
		if err.Error() == fmt.Sprintf("user with ID %d not found", task.UserID) {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("User with ID %d not found", task.UserID)})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating task"})
	}

	// Возвращаем только нужные поля
	response := map[string]interface{}{
		"id":      createdTask.ID,
		"task":    createdTask.Task,
		"is_done": createdTask.IsDone,
		"user_id": createdTask.UserID, // Если нужно
	}

	// Возвращаем ответ
	return c.JSON(http.StatusCreated, response)
}

// PatchTasksId - обновляет задачу по ID
func (h *TaskHandler) PatchTasksId(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid task ID"})
	}

	var task taskService.Message
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	updatedTask, err := h.Service.UpdateMessageByID(uint(taskID), task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Возвращаем обновленную задачу, без ненужных данных пользователя
	response := map[string]interface{}{
		"id":      updatedTask.ID,
		"task":    updatedTask.Task,
		"is_done": updatedTask.IsDone,
		"user_id": updatedTask.UserID,
	}

	return c.JSON(http.StatusOK, response)
}

// DeleteTasksId - удаляет задачу по ID
func (h *TaskHandler) DeleteTasksId(c echo.Context) error {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid task ID"})
	}

	err = h.Service.DeleteMessageByID(uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting task"})
	}

	return c.NoContent(http.StatusNoContent)
}
