package route

import (
	"github.com/Dpyde/Omchu/internal/chat"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupChatRoutes(app *fiber.App, db *gorm.DB) {
	chatRepo := chat.NewGormChatRepository(db)
	chatService := chat.NewChatService(chatRepo)
	chatHandler := chat.NewHttpChatHandler(chatService)

	chatRoutes := app.Group("/chat")
	// chatRoutes.Post("/", chatHandler.CreateChat)
	chatRoutes.Get("/:id", chatHandler.GetChat)
}
