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
	"github.com/Dpyde/Omchu/route"
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

	// userRepo := userGormRep.NewGormUserRepository(db)
	// userService := userSer.NewUserService(userRepo)
	// userHandler := userHndl.NewHttpUserHandler(userService)

	// authRepo := authGormRep.NewGormAuthRepository(db)
	// authService := authSer.NewAuthService(authRepo)
	// authHandler := authHndl.NewHttpAuthHandler(authService)

	// // Define routes
	// app.Post("/login", authHandler.Login)
	// app.Post("/register", authHandler.Register)

	// app.Post("/user", middleware.AuthMiddleware, userHandler.CreateUser)
	// app.Put("/user/:id", middleware.AuthMiddleware, userHandler.UpdateUser)
	// app.Delete("/user/:id", middleware.AuthMiddleware, userHandler.RemoveUser)
	// app.Get("/user/:id", middleware.AuthMiddleware, userHandler.FindByID)
	// app.Get("/users/:id", middleware.AuthMiddleware, userHandler.FindUsersToSwipe)
	route.SetupUserRoutes(app, db)
	route.SetupAuthRoutes(app, db)

	// Start the server
	app.Listen(":8000")
}
