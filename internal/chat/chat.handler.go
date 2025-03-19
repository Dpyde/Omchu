package chat

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type HttpChatHandler struct {
	service ChatService
}

func NewHttpChatHandler(service ChatService) *HttpChatHandler {
	return &HttpChatHandler{service: service}
}

func (h *HttpChatHandler) GetChat(c *fiber.Ctx) error {
	userId, ok := c.Locals("UserId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "UserId not found in context",
		})
	}
	chats, err := h.service.GetChat(userId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sucess": false,
			"error":  err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    chats,
	})
}
