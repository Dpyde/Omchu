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
	userRoutes.Post("/", middleware.Middleware, userHandler.CreateUser)
	userRoutes.Put("/", middleware.Middleware, userHandler.UpdateUser)
	userRoutes.Delete("/:id", middleware.Middleware, userHandler.RemoveUser)
	userRoutes.Get("/:id", middleware.Middleware, userHandler.FindByID)
	userRoutes.Get("/", middleware.Middleware, userHandler.GetMe)
	userRoutes.Get("/swipe/:id", middleware.Middleware, userHandler.FindUsersToSwipe)
}
