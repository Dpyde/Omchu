package database

import (
	"fmt"
	"log"
	"os"
	"time"

	auth "github.com/Dpyde/Omchu/internal/auth"
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
	db.Migrator().DropTable(&entity.User{}, &entity.Chat{}, &entity.Message{}, &entity.Picture{}, &entity.Swipe{})
	db.Migrator().DropTable("chat_users")
	//above is to clean the table before migration
	// Migrate the schema
	err = db.AutoMigrate(&entity.User{}, &entity.Chat{}, &entity.Message{}, &entity.Picture{}, &entity.Swipe{})
	if err != nil {
		return nil, err
	}

	// initialize the database with some data
	//newMessage := entity.Message{SenderID: 1,Text: "Hello World!"}

	// newSwipe := entity.Swipe{SwiperID: 1, SwipedID: 2, Liked: true}
	// // newSwipe := entity.Swipe{SwiperID: 1, Liked: true}
	newUser1 := entity.User{Name: "John Doe", Age: 12, Email: "myemail1", Color: "blue", Password: "password1"}
	// newUser2 := entity.User{Name: "Jane Doe", Age: 18, Email: "myemail2", Color: "blue", Password: "password2"}
	// newUser3 := entity.User{Name: "Juhn Doe", Age: 144, Email: "myemail3", Color: "blue", Password: "password3"}
	// newUser4 := entity.User{Name: "Jahn Doe", Age: 14, Email: "myemail4", Color: "blue", Password: "password4", Swipes: []entity.Swipe{newSwipe}}

	password1, err := auth.HashPassword("password")
	newUser2 := entity.User{Name: "Jane Doe", Age: 18, Email: "myemail2", Color: "blue", Password: "password2"}
	newSwipe := entity.Swipe{SwiperID: 1, SwipedID: 2, Liked: true}
	newUser := entity.User{Name: "John Doe", Age: 12, Email: "myemail", Color: "blue", Password: password1, Swipes: []entity.Swipe{newSwipe}}
	// db.Create(&newUser1)
	db.Create(&newUser)

	// db.Create(&newUser2)
	// db.Create(&newUser3)
	// db.Create(&newUser4)
	newChat := entity.Chat{Users: []entity.User{newUser1, newUser2}, Messages: []entity.Message{}}
	db.Create(&newChat)
	fmt.Println(newSwipe)

	// newChat := entity.Chat{Users: []entity.User{newUser,newUser2}, Messages: []entity.Message{newMessage}}
	// db.Create(&newChat)s
	return
}
