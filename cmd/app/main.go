package main

import (
	"FirstJobProject/internal/database"
	"FirstJobProject/internal/handlers"
	"FirstJobProject/internal/taskService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Автоматическая миграция модели Message
	if err := database.DB.AutoMigrate(&taskService.Message{}); err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	// Создание репозитория, сервиса и хендлера
	repo := taskService.NewMessageRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Создание маршрутизатора и добавление маршрутов
	router := mux.NewRouter()
	router.HandleFunc("/tasks", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/tasks", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/tasks/{id:[0-9]+}", handler.PatchMessageHandler).Methods("PATCH")
	router.HandleFunc("/tasks/{id:[0-9]+}", handler.DeleteMessageHandler).Methods("DELETE")

	// Запуск сервера
	http.ListenAndServe(":8080", router)
}
