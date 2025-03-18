package main

import (
	"fmt"
	"log"

	"github.com/Dpyde/Omchu/database"
	"github.com/Dpyde/Omchu/picture"
	"github.com/Dpyde/Omchu/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	if err1 := godotenv.Load(); err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Database connected")
	picture.InitR2()
	fmt.Println("Cloud connect")
	// Configure your PostgreSQL database details here

	route.SetupPictureRoutes(app, db)
	route.SetupChatRoutes(app, db)
	route.SetupUserRoutes(app, db)
	route.SetupAuthRoutes(app, db)
	route.SetupSwipeRoutes(app, db)

	app.Listen(":8000")
}
