package handlers

import (
	"FirstJobProject/internal/taskService"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *taskService.MessageService
}

func NewHandler(service *taskService.MessageService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message taskService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверяем, что поле "task" передано
	if message.Task == "" {
		http.Error(w, "Message field 'task' is required", http.StatusBadRequest)
		return
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

func (h *Handler) PatchMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL с помощью mux.Vars
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Message ID", http.StatusBadRequest)
		return
	}

	var message taskService.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Логируем полученные данные
	log.Printf("Received message: %+v", message)

	// Проверка на пустое поле task
	if message.Task == "" && message.IsDone == false {
		http.Error(w, "At least one field (task or is_done) must be provided", http.StatusBadRequest)
		return
	}

	// Обновляем сообщение
	updatedMessage, err := h.Service.UpdateMessageByID(uint(id), message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // Если не нашли ID, возвращаем 404
		return
	}

	// Отправляем обновленное сообщение с полями, которые были обновлены
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMessage)
}

func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL с помощью mux.Vars
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Message ID", http.StatusBadRequest)
		return
	}

	// Пытаемся удалить сообщение по ID
	err = h.Service.DeleteMessageByID(uint(id))
	if err != nil {
		// Если сообщение не найдено, возвращаем 404 Not Found
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Если удаление прошло успешно, возвращаем статус 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
