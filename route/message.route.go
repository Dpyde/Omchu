package route

import (
	"github.com/Dpyde/Omchu/internal/message"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUpMessageRoute(app *fiber.App, db *gorm.DB) {
	messageRepo := message.NewGormMessageRepository(db)
	messageService := message.NewMessageService(messageRepo)
	messageHandler := message.NewHttpMessageHandler(messageService)

	messageRoutes := app.Group("/message")
	messageRoutes.Get("/:chatId", messageHandler.GetMessage)
	messageRoutes.Post("/", messageHandler.SendMessage)
}
