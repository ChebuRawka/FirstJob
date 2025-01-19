package taskService

// MessageService - структура для работы с сообщениями
type MessageService struct {
	repo MessageRepository
}

// NewService - конструктор для создания нового сервиса
func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// CreateMessage - создание сообщения через сервис
func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

// GetAllMessages - получение всех сообщений через сервис
func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

// UpdateMessageByID - обновление сообщения через сервис
func (s *MessageService) UpdateMessageByID(id uint, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, message)
}

// DeleteMessageByID - удаление сообщения через сервис
func (s *MessageService) DeleteMessageByID(id uint) error {
	return s.repo.DeleteMessageByID(id)
}
