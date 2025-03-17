package route

import (
	authGormRep "github.com/Dpyde/Omchu/adapter/gorm/auth"
	authHndl "github.com/Dpyde/Omchu/adapter/http/auth"
	authSer "github.com/Dpyde/Omchu/internal/service/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupAuthRoutes(app *fiber.App, db *gorm.DB) {
	authRepo := authGormRep.NewGormAuthRepository(db)
	authService := authSer.NewAuthService(authRepo)
	authHandler := authHndl.NewHttpAuthHandler(authService)

	authRoutes := app.Group("/auth")
	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/register", authHandler.Register)
}
