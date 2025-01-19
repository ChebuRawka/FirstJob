package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//var task string

// Структура для тела запроса
type requestBody struct {
	Message interface{} `json:"message"` // Поле для обработки данных из тела запроса
}

// Структура для ответа
type responseBody struct {
	Message string `json:"message"`         // Ответ с сообщением
	Error   string `json:"error,omitempty"` // поле для ошибок
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var body requestBody // Создаем переменную body для хранения данных из тела запроса

	// Проверка на пустое тело
	if r.Body == nil || r.ContentLength == 0 { // если тело отсутствует или если длина содержимого тела равно 0
		http.Error(w, "Request body is empty", http.StatusBadRequest) // отправляем ошибку
		return
	}

	// Декодируем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil { // если произошла ошибка при декодировании, мы выводим ошибку в лог и отправляем ответ
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Проверяем тип данных поля message
	if message, ok := body.Message.(string); !ok {
		http.Error(w, "Invalid data type for message", http.StatusBadRequest)
		return
	} else {
		// Если тип данных правильный, можно продолжать
		log.Printf("Message: %s", message)
	}
	// создаем объект Message (используя структуру из orm.go)
	message := Message{
		Task:   body.Message.(string), // присваиваем сообщение из запроса
		IsDone: false,                 // можно установить значение по умолчанию
	}

	//  Сохраняем в БД
	if err := DB.Create(&message).Error; err != nil {
		log.Printf("Error saving message to DB:", err)
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	// Обновляем глобальную переменную task
	//	task = body.Message // если все декодирование прошло успешно, мы присваиваем значение поля Message из структуры body глобал переменной task

	// Формируем ответ
	response := responseBody{ // создаем объект response responseBody
		Message: fmt.Sprintf("Task created: %s", message.Task),
	}

	// Отправляем успешный ответ
	w.Header().Set("Content-Type", "application/json") // тело ответа будет содержать данные в формате JSON
	w.WriteHeader(http.StatusOK)                       // отправляем статус HTTP 200 (OK), который означает успешный обработанный запрос.
	json.NewEncoder(w).Encode(response)                // кодирования объекта response в JSON и отправки его в тело ответа.
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message

	//извлекаем все сообщения из бд с помощью DB.Find
	if err := DB.Find(&messages).Error; err != nil {
		log.Printf("Error fetching messages from database: %v", err)
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	//создание ответа
	//response := responseBody{ // создаем объект responseBody
	//	Message: fmt.Sprintf("Hello, %s!", task), // значение поля Message формируется с помощью Sprintf
	//}

	//возвращаем все сообщения в JSON формате
	w.Header().Set("Content-Type", "application/json") // тело ответа будет содержать данные в формате JSON
	json.NewEncoder(w).Encode(messages)                // отправка json ответа. Encode - функция кодирует объект в json формат и отправляет в тело ответа
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	// Создает таблицу в бд для модели Message
	if err := DB.AutoMigrate(&Message{}); err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	//маршрутизация
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")

	http.ListenAndServe(":8080", router)
}
