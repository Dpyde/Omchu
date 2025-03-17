package route

import (
	chatGormRep "github.com/Dpyde/Omchu/adapter/gorm/chat"
	chatHndl "github.com/Dpyde/Omchu/adapter/http/chat"
	chatSer "github.com/Dpyde/Omchu/internal/service/chat"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupChatRoutes(app *fiber.App, db *gorm.DB) {
	chatRepo := chatGormRep.NewGormChatRepository(db)
	chatService := chatSer.NewChatService(chatRepo)
	chatHandler := chatHndl.NewHttpChatHandler(chatService)

	chatRoutes := app.Group("/chat")
	// chatRoutes.Post("/", chatHandler.CreateChat)
	chatRoutes.Get("/:id", chatHandler.GetChat)
}
