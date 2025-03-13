package chatSer

import (
	"github.com/Dpyde/Omchu/internal/entity"
	chatRep "github.com/Dpyde/Omchu/internal/repository/chat"
)

type ChatService interface {
	GetChat(userId string) ([]entity.Chat, error)
}

type chatServiceImpl struct {
	repo chatRep.ChatRepository
}

func NewChatService(repo chatRep.ChatRepository) ChatService {
	return &chatServiceImpl{repo: repo}
}

func (s *chatServiceImpl) GetChat(userId string) ([]entity.Chat, error) {
	chat, err := s.repo.FindById(userId)
	if err != nil {
		return []entity.Chat{}, err
	}
	return chat, nil
}
