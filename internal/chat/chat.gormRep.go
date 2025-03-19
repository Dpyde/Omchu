package chat

import (
	"fmt"
	"strconv"

	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/gorm"
)

type GormChatRepository struct {
	db *gorm.DB
}

type ExtendedChat struct {
	ChatID uint
	Username string
	PictureProfile entity.Picture
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
		fmt.Print(userID)
		if err != nil {
			return []ExtendedChat{}, err
		}
		var user entity.User
		if err := r.db.Preload("Pictures").First(&user, userIDInt).Error; err != nil {
			return nil, err
		}
		noti := latestMessage.ID != 0 && !latestMessage.Read && int(latestMessage.SenderID) != userIDInt // Check if the latest message is unread
		Pictures := user.Pictures
		var profilePic entity.Picture 
		if len(Pictures) > 0 {
			profilePic = Pictures[0]
		}
		extendedChats = append(extendedChats, ExtendedChat{
			ChatID: chat.ID,
			Username: user.Name,
			PictureProfile: profilePic,
			Noti:   noti,
		})
	}

	return extendedChats, nil
}
