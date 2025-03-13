package main

import (
	"fmt"
	"log"

	chatGormRep "github.com/Dpyde/Omchu/adapter/gorm/chat"
	userGormRep "github.com/Dpyde/Omchu/adapter/gorm/user"
	chatHndl "github.com/Dpyde/Omchu/adapter/http/chat"
	userHndl "github.com/Dpyde/Omchu/adapter/http/user"
	"github.com/Dpyde/Omchu/database"
	chatSer "github.com/Dpyde/Omchu/internal/service/chat"
	userSer "github.com/Dpyde/Omchu/internal/service/user"
	"github.com/gofiber/fiber/v2"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	app := fiber.New()
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")

	userRepo := userGormRep.NewGormUserRepository(db)
	userService := userSer.NewUserService(userRepo)
	userHandler := userHndl.NewHttpUserHandler(userService)

	chatRepo := chatGormRep.NewGormChatRepository(db)
	chatService := chatSer.NewChatService(chatRepo)
	chatHandler := chatHndl.NewHttpChatHandler(chatService)
	// Define routes

	app.Post("/user", userHandler.CreateUser)
	app.Get("/chat/:userid", chatHandler.GetChat)
	// Start the server
	app.Listen(":8000")
}
