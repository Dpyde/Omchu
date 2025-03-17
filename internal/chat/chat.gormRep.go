package chat

import (
	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/gorm"
)

type GormChatRepository struct {
	db *gorm.DB
}

type ExtendedChat struct {
	Chat entity.Chat
	Noti bool
}

func NewGormChatRepository(db *gorm.DB) ChatRepository {
	return &GormChatRepository{db: db}
}

func (r *GormChatRepository) FindById(userId string) ([]ExtendedChat, error) {
	var chats []entity.Chat
	err := r.db.Joins("JOIN chat_users cu ON cu.chat_id = chats.id").
		Where("cu.user_id = ?", userId).
		Find(&chats).Error
	if err != nil {
		return []ExtendedChat{}, err
	}

	var extendedChats []ExtendedChat
	for _, chat := range chats {
		var latestMessage entity.Message
		err := r.db.Where("chat_id = ?", chat.ID).Order("created_at DESC").First(&latestMessage).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return []ExtendedChat{}, err
		}

		noti := !latestMessage.Read // Check if the latest message is unread
		extendedChats = append(extendedChats, ExtendedChat{Chat: chat, Noti: noti})
	}

	return extendedChats, nil
}
