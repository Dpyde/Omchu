package chatRep

import (
	"github.com/Dpyde/Omchu/internal/entity"
)

type ChatRepository interface {
	FindById(userId string) ([]entity.Chat, error)
}