package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      uint
	Email    string `gorm:"unique"`
	Color    string
	Password string
	Pictures []Picture
	Chats    []Chat  `gorm:"many2many:chat_users;"`
	Swipes   []Swipe `gorm:"foreignKey:SwiperID"`
}

type Chat struct {
	gorm.Model
	Users    []User `gorm:"many2many:chat_users;"`
	Messages []Message
}

type Message struct {
	gorm.Model
	SenderID uint
	ChatID   uint
	Text     string
	Read     bool `gorm:"default:false"`
}

type Picture struct {
	UserID uint
	Url    string
	Key    string
}

type Swipe struct {
	gorm.Model
	SwiperID uint `gorm:"primaryKey"`
	SwipedID uint `gorm:"primaryKey"`
	Liked    bool
}
