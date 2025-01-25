package handlers

import (
	"FirstJobProject/internal/taskService"
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

	return c.JSON(http.StatusOK, allTasks)
}

// PostTasks - создает новую задачу
func (h *TaskHandler) PostTasks(c echo.Context) error {
	var task taskService.Message
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	createdTask, err := h.Service.CreateMessage(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating task"})
	}

	return c.JSON(http.StatusCreated, createdTask)
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error updating task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
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
