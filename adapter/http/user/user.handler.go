package userHndl

import (
	"fmt"

	authHndl "github.com/Dpyde/Omchu/adapter/http/auth"
	"github.com/Dpyde/Omchu/internal/entity"
	userSer "github.com/Dpyde/Omchu/internal/service/user"
	"github.com/gofiber/fiber/v2"
)

// Primary adapter
type HttpUserHandler struct {
	service userSer.UserService
}

func NewHttpUserHandler(service userSer.UserService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateUser(user); err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	authHndl.SendTokenResponse(c, user, fiber.StatusAccepted)
	return c.Status(fiber.StatusCreated).JSON(user)
}
