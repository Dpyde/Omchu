package swipe

import (

	"strconv"

	"github.com/Dpyde/Omchu/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type HtttSwipeHandler struct {
	service SwipeService
}

func NewHttpSwipeHandler(service SwipeService) *HtttSwipeHandler {
	return &HtttSwipeHandler{service: service}
}

func (h *HtttSwipeHandler) SwipeCheck(c *fiber.Ctx) error {
	var swipe entity.Swipe
	if err := c.BodyParser(&swipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success":false , "error": err.Error()})
	}
	var is_match bool
	userIdStr,ok := c.Locals("UserId").(string)
	if(!ok) {return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success":false,"error": "No UserId in local"})}
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success":false,"error": "Invalid UserId in local"})
	}
	swipe.SwiperID = uint(userId)
	if err := h.service.SwipeCheck(&swipe, &is_match); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success":false,"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"data":   swipe,
		"is_match": is_match,
	})
}