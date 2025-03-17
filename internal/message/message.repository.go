package message

import "github.com/Dpyde/Omchu/internal/entity"

type MessageRepository interface {
	GetMessage(chatId string) ([]entity.Message, error)
	SendMessage(message *entity.Message) error
}
