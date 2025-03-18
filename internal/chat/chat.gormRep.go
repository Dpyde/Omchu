package chat

import (
	"strconv"

	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/gorm"
)

type GormChatRepository struct {
	db *gorm.DB
}

type ExtendedChat struct {
	ChatID uint
	UserID uint
	Noti   bool
}

func NewGormChatRepository(db *gorm.DB) ChatRepository {
	return &GormChatRepository{db: db}
}

func (r *GormChatRepository) FindById(userId string) ([]ExtendedChat, error) {
	var chats []entity.Chat
	err := r.db.Joins("JOIN chat_users cu ON cu.chat_id = chats.id").
		Where("cu.user_id = ?", userId).
		Preload("Users").
		Find(&chats).Error
	if err != nil {
		return []ExtendedChat{}, err
	}

	var extendedChats []ExtendedChat
	for _, chat := range chats {
		var userID uint
		for _, user := range chat.Users {
			intid, err := strconv.Atoi(userId)
			if condition := err != nil; condition {
				return []ExtendedChat{}, err
			}
			if int(user.ID) != intid {
				userID = user.ID
			}
		}
		var latestMessage entity.Message
		msgErr := r.db.Where("chat_id = ?", chat.ID).
			Order("created_at DESC").
			First(&latestMessage).Error
		if msgErr != nil && msgErr != gorm.ErrRecordNotFound {
			continue // Skip this chat if an error occurs
		}

		userIDInt, err := strconv.Atoi(userId)
		if err != nil {
			return []ExtendedChat{}, err
		}
		noti := latestMessage.ID != 0 && !latestMessage.Read && int(latestMessage.SenderID) != userIDInt // Check if the latest message is unread
		extendedChats = append(extendedChats, ExtendedChat{
			ChatID: chat.ID,
			UserID: userID,
			Noti:   noti,
		})
	}

	return extendedChats, nil
}
