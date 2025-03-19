package user

import (
	"fmt"
	"strconv"

	"github.com/Dpyde/Omchu/internal/entity"
	"github.com/gofiber/fiber/v2"
)

// Primary adapter
type HttpUserHandler struct {
	service UserService
}

func NewHttpUserHandler(service UserService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user entity.User
	// fmt.Println("checkPoint1")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}

	newUser, err := h.service.CreateUser(user)

	// fmt.Println("checkPoint2")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "user": newUser})
}

func (h *HttpUserHandler) FindUsersToSwipe(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid user ID"})
	}
	users, err := h.service.FindUsersToSwipe(uint(id))
	// fmt.Println("checkPoint1")
	if err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "users": users})
}

func (h *HttpUserHandler) UpdateUser(c *fiber.Ctx) error {
	var newUser entity.User
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		// fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid user ID"})
	}
	if err := c.BodyParser(&newUser); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}
	// newUser.Pictures = UploadPicsToR2()
	updatedUser, err := h.service.UpdateUser(newUser, uint(id))
	if err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "user": updatedUser})
}

func (h *HttpUserHandler) FindByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid user ID"})
	}
	user, err := h.service.FindByID(uint(id))
	if err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "user": user})
}

func (h *HttpUserHandler) RemoveUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid user ID"})
	}
	fmt.Println("removePoint1")
	err = h.service.RemoveUser(uint(id))
	if err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "user deleted successfully"})

}

// func (h *HttpUserHandler) FindByUsername(c *fiber.Ctx) error {
// 	username := c.Params("username")
// 	user, err := h.service.FindByUsername(username)
// 	if err != nil {
// 		// Return an appropriate error message and status code
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "user": user})
// }

// func (h *HttpUserHandler) FindByEmail(c *fiber.Ctx) error {
// 	email := c.Params("email")
// 	user, err := h.service.FindByEmail(email)
// 	if err != nil {
// 		// Return an appropriate error message and status code
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "user": user})
// }
