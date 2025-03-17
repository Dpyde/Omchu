package message

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type MessageService interface {
	GetMessage(chatId string, UserId string) ([]entity.Message, error)
	SendMessage(message *entity.Message) error
}

type messageServiceImpl struct {
	repo MessageRepository
}

func NewMessageService(repo MessageRepository) MessageService {
	return &messageServiceImpl{repo: repo}
}

func (s *messageServiceImpl) GetMessage(chatId string, userId string) ([]entity.Message, error) {
	messages, err := s.repo.GetMessage(chatId, userId)
	if err != nil {
		return []entity.Message{}, err
	}
	return messages, nil
}

func (s *messageServiceImpl) SendMessage(message *entity.Message) error {
	if err := s.repo.SendMessage(message); err != nil {
		return err
	}
	return nil
}
