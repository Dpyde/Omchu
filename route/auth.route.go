package route

import (
	"github.com/Dpyde/Omchu/internal/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupAuthRoutes(app *fiber.App, db *gorm.DB) {
	authRepo := auth.NewGormAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewHttpAuthHandler(authService)

	authRoutes := app.Group("/auth")
	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/register", authHandler.Register)
}
