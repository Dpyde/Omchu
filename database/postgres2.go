package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Dpyde/Omchu/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() (db *gorm.DB, err error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "myuser"
		password = "mypassword"
		dbname   = "mydatabase"
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// Drop tables before migration (for development purposes only)
	db.Migrator().DropTable(
		&entity.User{}, &entity.Chat{}, &entity.Message{}, &entity.Picture{}, &entity.Swipe{}, "chat_users",
	)

	// Migrate the schema
	err = db.AutoMigrate(
		&entity.User{}, &entity.Chat{}, &entity.Message{}, &entity.Picture{}, &entity.Swipe{},
	)
	if err != nil {
		return nil, err
	}

	// Initialize the database with more test data
	users := []entity.User{
		{Name: "Alice Smith", Age: 25, Email: "alice@example.com", Color: "red", Password: "alicepass"},
		{Name: "Bob Johnson", Age: 30, Email: "bob@example.com", Color: "green", Password: "bobpass"},
		{Name: "Charlie Brown", Age: 28, Email: "charlie@example.com", Color: "yellow", Password: "charliepass"},
	}
	db.Create(&users)

	swipes := []entity.Swipe{
		{SwiperID: 1, SwipedID: 2, Liked: true},
		{SwiperID: 2, SwipedID: 3, Liked: false},
		{SwiperID: 3, SwipedID: 1, Liked: true},
	}
	db.Create(&swipes)

	pictures := []entity.Picture{
		{UserID: 1, Url: "https://example.com/alice1.jpg", Key: "alice1"},
		{UserID: 2, Url: "https://example.com/bob1.jpg", Key: "bob1"},
		{UserID: 3, Url: "https://example.com/charlie1.jpg", Key: "charlie1"},
	}
	db.Create(&pictures)

	newChat := entity.Chat{Users: []entity.User{users[0], users[1]}, Messages: []entity.Message{
		{SenderID: 1, ChatID: 1, Text: "Hey Bob!", Read: true},
		{SenderID: 2, ChatID: 1, Text: "Hey Alice!", Read: false},
	}}
	db.Create(&newChat)

	return db, nil
}