package userHndl

import (
	"fmt"

	"strconv"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}

	if err := h.service.CreateUser(user); err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "user": user})
}

func (h *HttpUserHandler) UpdateUser(c *fiber.Ctx) error {
	var newUser entity.User
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid user ID"})
	}
	if err := c.BodyParser(&newUser); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}
	updatedUser, err := h.service.UpdateUser(newUser, uint(id))
	if err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "user": updatedUser})
}
