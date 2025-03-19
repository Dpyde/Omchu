package main

import (
	"fmt"
	"log"

	"github.com/Dpyde/Omchu/database"
	"github.com/Dpyde/Omchu/picture"
	"github.com/Dpyde/Omchu/hubrouter"
	"github.com/Dpyde/Omchu/internal/ws"
	"github.com/Dpyde/Omchu/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Accepts requests from any origin (IP)
		AllowMethods: "GET,POST,PUT,DELETE", // Allowed HTTP methods
		AllowHeaders: "Content-Type,Authorization", // Allowed headers
	}))
	
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	if err1 := godotenv.Load(); err1 != nil {
		log.Fatal(err1)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	route.SetupMessageRoute(app, db)
	go app.Listen(":8000")

	hub := ws.NewHub()
	WsHandler := ws.NewHandler(hub)
	go hub.Run(db)

	hubrouter.InitRouter(WsHandler)
	// Start the server
	hubrouter.Start("127.0.0.1:8080")
}
