package main

import (
	"fmt"
	"log"

	userGormRep "github.com/Dpyde/Omchu/adapter/gorm/user"
	userHndl "github.com/Dpyde/Omchu/adapter/http/user"
	"github.com/Dpyde/Omchu/database"
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

	// Define routes
	app.Post("/user", userHandler.CreateUser)

	// Start the server
	app.Listen(":8000")
}
