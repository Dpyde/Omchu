package main

import (
	"fmt"
	"log"

	// authGormRep "github.com/Dpyde/Omchu/adapter/gorm/auth"
	// userGormRep "github.com/Dpyde/Omchu/adapter/gorm/user"
	// authHndl "github.com/Dpyde/Omchu/adapter/http/auth"
	// userHndl "github.com/Dpyde/Omchu/adapter/http/user"
	// "github.com/Dpyde/Omchu/database"
	// authSer "github.com/Dpyde/Omchu/internal/service/auth"
	// userSer "github.com/Dpyde/Omchu/internal/service/user"
	// middleware "github.com/Dpyde/Omchu/middleware"
	"github.com/Dpyde/Omchu/database"
	"github.com/Dpyde/Omchu/internal/hubrouter"
	"github.com/Dpyde/Omchu/internal/ws"
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
	fmt.Println("Database connected")
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Configure your PostgreSQL database details here

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
