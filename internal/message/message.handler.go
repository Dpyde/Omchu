package message

import (
	"github.com/Dpyde/Omchu/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type HttpMessageHandler struct {
	service MessageService
}

func NewHttpMessageHandler(service MessageService) *HttpMessageHandler {
	return &HttpMessageHandler{service: service}
}

func (h *HttpMessageHandler) GetMessage(c *fiber.Ctx) error {
	chatId := c.Params("chatId")
	UserId, ok := c.Locals("UserId").(string)
	if !ok {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"message": "UserId not found in context",
	})

	messages, err := h.service.GetMessage(chatId, UserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    messages,
	})
}

func (h *HttpMessageHandler) SendMessage(c *fiber.Ctx) error {
	var message entity.Message
	if err := c.BodyParser(&message); err != nil {

	}

	if err := h.service.SendMessage(&message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Message created",
	})
}
