package swipeHndl

import (
	"github.com/Dpyde/Omchu/internal/entity"
	"github.com/Dpyde/Omchu/internal/service/swipe"
	"github.com/gofiber/fiber/v2"
)

type HtttSwipeHandler struct {
	service swipeSer.SwipeService
}

func NewHttpSwipeHandler(service swipeSer.SwipeService) *HtttSwipeHandler {
	return &HtttSwipeHandler{service: service}
}

func (h *HtttSwipeHandler) SwipeCheck(c *fiber.Ctx) error {
	var swipe entity.Swipe
	if err := c.BodyParser(&swipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var is_match bool
	if err := h.service.SwipeCheck(&swipe, &is_match); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"data":   swipe,
		"is_match": is_match,
	})
}

