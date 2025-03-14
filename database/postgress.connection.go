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
		host     = "localhost"  // or the Docker service name if running in another container
		port     = 5432         // default PostgreSQL port
		user     = "myuser"     // as defined in docker-compose.yml
		password = "mypassword" // as defined in docker-compose.yml
		dbname   = "mydatabase" // as defined in docker-compose.yml
	)
	//to config later
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger configuration
	})
	if err != nil {
		return nil, err
	}

	db.Migrator().DropTable(&entity.User{}, &entity.Chat{}, &entity.Message{}, &entity.Picture{}, &entity.Swipe{})
	db.Migrator().DropTable("chat_users")
	//above is to clean the table before migration
	// Migrate the schema
	err = db.AutoMigrate(&entity.User{}, &entity.Chat{}, &entity.Message{}, &entity.Picture{}, &entity.Swipe{})
	if err != nil {
		return nil, err
	}

	// initialize the database with some data
	newMessage := entity.Message{SenderID: 1, Text: "Hello World!"}
	// newUser := entity.User{Name: "John Doe", Age: 12, Email: "myemail", Color: "blue", PhoneNumber: "1234567890", Password: "password"}
	// newUser2 := entity.User{Name: "Jane Doe", Age: 18, Email: "myemail2", Color: "blue", PhoneNumber: "2222222222", Password: "password2"}
	newUser := entity.User{Name: "John Doe", Age: 12, Email: "myemail", Color: "blue", Password: "password"}
	newUser2 := entity.User{Name: "Jane Doe", Age: 18, Email: "myemail2", Color: "blue", Password: "password2"}
	db.Create(&newUser)
	db.Create(&newUser2)
	newChat := entity.Chat{Users: []entity.User{newUser, newUser2}, Messages: []entity.Message{newMessage}}
	db.Create(&newChat)
	return
}
