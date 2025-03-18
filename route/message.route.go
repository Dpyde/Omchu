package route

import (
	"github.com/Dpyde/Omchu/internal/message"
	"github.com/Dpyde/Omchu/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupMessageRoute(app *fiber.App, db *gorm.DB) {
	messageRepo := message.NewGormMessageRepository(db)
	messageService := message.NewMessageService(messageRepo)
	messageHandler := message.NewHttpMessageHandler(messageService)

	messageRoutes := app.Group("/message")
	messageRoutes.Use(middleware.Middleware)
	messageRoutes.Post("/", messageHandler.SendMessage)
	messageRoutes.Get("/:chatId", messageHandler.GetMessage)

}
