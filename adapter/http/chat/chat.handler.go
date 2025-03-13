package chatHandlr

import (
	"fmt"

	chatSer "github.com/Dpyde/Omchu/internal/service/chat"
	"github.com/gofiber/fiber/v2"
)

type HttpChatHandler struct {
	service chatSer.ChatService
}

func NewHttpChatHandler(service chatSer.ChatService) *HttpChatHandler {
	return &HttpChatHandler{service: service}
}

func (h *HttpChatHandler) GetChat(c *fiber.Ctx) error {
	userId := c.Params("userId")
	chat, err := h.service.GetChat(userId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(chat)
}
