package route

import (
	"github.com/Dpyde/Omchu/internal/chat"
	middleware "github.com/Dpyde/Omchu/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupChatRoutes(app *fiber.App, db *gorm.DB) {
	chatRepo := chat.NewGormChatRepository(db)
	chatService := chat.NewChatService(chatRepo)
	chatHandler := chat.NewHttpChatHandler(chatService)

	chatRoutes := app.Group("/chat")
	chatRoutes.Use(middleware.Middleware)
	chatRoutes.Get("/", chatHandler.GetChat)
}
