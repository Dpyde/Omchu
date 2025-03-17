package chatGormRep

import (
	"github.com/Dpyde/Omchu/internal/entity"
	chatRep "github.com/Dpyde/Omchu/internal/repository/chat"
	"gorm.io/gorm"
)

type GormChatRepository struct {
	db *gorm.DB
}

func NewGormChatRepository(db *gorm.DB) chatRep.ChatRepository {
	return &GormChatRepository{db: db}
}

func (r *GormChatRepository) FindById(userId string) ([]entity.Chat, error) {
	var chats []entity.Chat
	err := r.db.Joins("JOIN chat_users cu ON cu.chat_id = chats.id").
		Where("cu.user_id = ?", userId).
		Find(&chats).Error
	if err != nil {
		return []entity.Chat{}, err
	}
	return chats, err
}
