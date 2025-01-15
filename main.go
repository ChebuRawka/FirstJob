package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var task string

// Структура для тела запроса
type requestBody struct {
	Message string `json:"message"` // Поле для обработки данных из тела запроса
}

// Структура для ответа
type responseBody struct {
	Message string `json:"message"` // Ответ с сообщением
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
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

	// Обновляем глобальную переменную task
	task = body.Message // если все декодирование прошло успешно, мы присваиваем значение поля Message из структуры body глобал переменной task

	// Формируем ответ
	response := responseBody{ // создаем объект response responseBody
		Message: fmt.Sprintf("Task updated to: %s", task),
	}

	// Отправляем успешный ответ
	w.Header().Set("Content-Type", "application/json") // тело ответа будет содержать данные в формате JSON
	w.WriteHeader(http.StatusOK)                       // отправляем статус HTTP 200 (OK), который означает успешный обработанный запрос.
	json.NewEncoder(w).Encode(response)                // кодирования объекта response в JSON и отправки его в тело ответа.
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	//создание ответа
	response := responseBody{ // создаем объект responseBody
		Message: fmt.Sprintf("Hello, %s!", task), // значение поля Message формируется с помощью Sprintf
	}
	w.Header().Set("Content-Type", "application/json") // тело ответа будет содержать данные в формате JSON
	json.NewEncoder(w).Encode(response)                // отправка json ответа. Encode - функция кодирует объект в json формат и отправляет в тело ответа
}

func main() {
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/get", GetHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
