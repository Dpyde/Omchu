package route

import (
	middleware "github.com/Dpyde/Omchu/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := user.NewGormUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewHttpUserHandler(userService)

	userRoutes := app.Group("/user")
	userRoutes.Post("/", middleware.AuthMiddleware, userHandler.CreateUser)
	userRoutes.Put("/:id", middleware.AuthMiddleware, userHandler.UpdateUser)
	userRoutes.Delete("/:id", middleware.AuthMiddleware, userHandler.RemoveUser)
	userRoutes.Get("/:id", middleware.AuthMiddleware, userHandler.FindByID)
	userRoutes.Get("/swipe/:id", middleware.AuthMiddleware, userHandler.FindUsersToSwipe)
}
