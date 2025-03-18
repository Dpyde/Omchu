package message

import (
	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/gorm"
)

type GormMessageRepository struct {
	db *gorm.DB
}

func NewGormMessageRepository(db *gorm.DB) MessageRepository {
	return &GormMessageRepository{db: db}
}

func (r *GormMessageRepository) GetMessage(chatId string, userId string) ([]entity.Message, error) {
	var messages []entity.Message
	err := r.db.Where("chatID = ?", chatId).Find(&messages).Error
	if err != nil {
		return []entity.Message{}, err
	}
	r.db.Model(&entity.Message{}).
		Where("chatID = ? AND userID != ? AND read = ?", chatId, userId, false).
		Update("read", true)
	return messages, nil
}

func (r *GormMessageRepository) SendMessage(message *entity.Message) error {
	err := r.db.Create(&message).Error
	if err != nil {
		return err
	}
	return nil
}
