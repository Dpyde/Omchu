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
	messages, err := h.service.GetMessage(chatId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(messages)
}

func (h *HttpMessageHandler) SendMessage(c *fiber.Ctx) error {
	var message entity.Message
	if err := c.BodyParser(&message); err != nil {

	}

	if err := h.service.SendMessage(&message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Message created"})
}
