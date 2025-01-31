package main

import (
	"FirstJobProject/internal/database"
	"FirstJobProject/internal/handlers"
	"FirstJobProject/internal/taskService"
	"FirstJobProject/internal/userService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Инициализируем базу данных
	database.InitDB()

	// Автоматическое применение миграций для задач
	if err := database.DB.AutoMigrate(&taskService.Message{}); err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	// Создаем репозиторий и сервис для задач
	tasksRepo := taskService.NewMessageRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	// Создаем репозиторий и сервис для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализируем Echo
	e := echo.New()

	// Используем Logger и Recover для обработки ошибок
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем маршруты для задач
	e.GET("/tasks", tasksHandler.GetTasks)
	e.POST("/tasks", tasksHandler.PostTasks)
	e.PATCH("/tasks/:id", tasksHandler.PatchTasksId)
	e.DELETE("/tasks/:id", tasksHandler.DeleteTasksId)

	//Маршрут для получения всех задач для конкретного пользака:

	e.GET("/users/:id/tasks", tasksHandler.GetTasksByUserID)

	// Регистрируем маршруты для пользователей
	e.GET("/users", userHandler.GetUsers)
	e.POST("/users", userHandler.PostUser)
	e.PATCH("/users/:id", userHandler.PatchUserByID)
	e.DELETE("/users/:id", userHandler.DeleteUserByID)

	// Запускаем сервер
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
