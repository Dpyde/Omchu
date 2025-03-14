package main

import (
	"fmt"
	"log"
  "github.com/Dpyde/Omchu/adapter/gorm/user"
  "github.com/Dpyde/Omchu/adapter/http/user"
  "github.com/Dpyde/Omchu/internal/service/user"
  "github.com/Dpyde/Omchu/database"
  "github.com/gofiber/fiber/v2"
  "github.com/Dpyde/Omchu/adapter/gorm/swipe"
  "github.com/Dpyde/Omchu/adapter/http/swipe"
  "github.com/Dpyde/Omchu/internal/service/swipe"
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
	db,err := database.InitDatabase();
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Database connected")

  userRepo := userGormRep.NewGormUserRepository(db)
  userService := userSer.NewUserService(userRepo)
  userHandler := userHndl.NewHttpUserHandler(userService)

  swipeRepo := swipeGormRep.NewGormSwipeRepository(db)
  swipeService := swipeSer.NewSwipeService(swipeRepo)
  swipeHandler := swipeHndl.NewHttpSwipeHandler(swipeService)


  // Define routes
  app.Post("/user", userHandler.CreateUser)
  app.Post("/swipe", swipeHandler.SwipeCheck)

  // Start the server
  app.Listen(":8000")
}
