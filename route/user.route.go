package route

import (
	"github.com/Dpyde/Omchu/internal/user"
	middleware "github.com/Dpyde/Omchu/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := user.NewGormUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewHttpUserHandler(userService)

	userRoutes := app.Group("/user")
	userRoutes.Use(middleware.Middleware)
	userRoutes.Post("/", userHandler.CreateUser)
	userRoutes.Put("/:id", userHandler.UpdateUser)
	userRoutes.Delete("/:id", userHandler.RemoveUser)
	userRoutes.Get("/:id", userHandler.FindByID)
	userRoutes.Get("/swipe/:id", userHandler.FindUsersToSwipe)
}
