package route

import (
	userGormRep "github.com/Dpyde/Omchu/adapter/gorm/user"
	userHndl "github.com/Dpyde/Omchu/adapter/http/user"
	userSer "github.com/Dpyde/Omchu/internal/service/user"
	middleware "github.com/Dpyde/Omchu/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := userGormRep.NewGormUserRepository(db)
	userService := userSer.NewUserService(userRepo)
	userHandler := userHndl.NewHttpUserHandler(userService)

	userRoutes := app.Group("/user")
	userRoutes.Post("/", middleware.AuthMiddleware, userHandler.CreateUser)
	userRoutes.Put("/:id", middleware.AuthMiddleware, userHandler.UpdateUser)
	userRoutes.Delete("/:id", middleware.AuthMiddleware, userHandler.RemoveUser)
	userRoutes.Get("/:id", middleware.AuthMiddleware, userHandler.FindByID)
	userRoutes.Get("/swipe/:id", middleware.AuthMiddleware, userHandler.FindUsersToSwipe)
}
