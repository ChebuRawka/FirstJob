package handlers

import (
	"FirstJobProject/internal/taskService"
	//"FirstJobProject/internal/utils"
	"FirstJobProject/internal/web/tasks"
	"context"
	"gorm.io/gorm"
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"log"
	//"net/http"
	//"strconv"
)

type Handler struct {
	Service *taskService.MessageService
}

func NewHandler(service *taskService.MessageService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := uint(request.Id)

	// Пытаемся удалить задачу по ID
	err := h.Service.DeleteMessageByID(taskID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return tasks.DeleteTasksId404JSONResponse{}, nil
		}
		return tasks.DeleteTasksId500JSONResponse{}, nil
	}

	// Если задача успешно удалена
	return tasks.DeleteTasksId204Response{}, nil
}
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := uint(request.Id)

	// Пытаемся получить текущую задачу
	currentTask, err := h.Service.GetMessageByID(taskID, taskService.Message{})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return tasks.PatchTasksId404JSONResponse{}, nil
		}
		return tasks.PatchTasksId500JSONResponse{}, nil
	}

	// Обновляем поля, только если они переданы
	taskRequest := request.Body
	if request.Body.Task != nil {
		currentTask.Task = *taskRequest.Task
	}
	if request.Body.IsDone != nil {
		currentTask.IsDone = *taskRequest.IsDone
	}

	// Пытаемся сохранить изменения
	updatedTask, err := h.Service.UpdateMessageByID(taskID, currentTask)
	if err != nil {
		return tasks.PatchTasksId500JSONResponse{}, nil
	}

	// Возвращаем обновленную задачу
	return tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}
func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Message{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateMessage(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}
