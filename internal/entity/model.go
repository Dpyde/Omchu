package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Age   uint
	Email string `gorm:"unique"`
	Color string
	// PhoneNumber string `gorm:"unique"`
	Password string
	Pictures []Picture
	Chats    []Chat  `gorm:"many2many:chat_users;"`
	swipes   []Swipe `gorm:"foreignKey:SwiperID"`
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
}

type Picture struct {
	UserID uint
	Url    string
	Key    string
}

type Swipe struct {
	gorm.Model
	SwiperID uint
	SwipedID uint
	Liked    bool
}
