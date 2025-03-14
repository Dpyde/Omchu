package authHndl

import (
	"os"
	"strconv"
	"time"

	"github.com/Dpyde/Omchu/internal/entity"
	authSer "github.com/Dpyde/Omchu/internal/service/auth"
	"github.com/gofiber/fiber/v2"
)

type HttpAuthHandler struct {
	service authSer.AuthService
}

func NewHttpAuthHandler(service authSer.AuthService) *HttpAuthHandler {
	return &HttpAuthHandler{service: service}
}

func (h *HttpAuthHandler) Login(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}

	if err := h.service.Login(user.Email, user.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	// return c.JSON(fiber.Map{"token": token})
	return SendTokenResponse(c, user, fiber.StatusOK)
}

func (h *HttpAuthHandler) Register(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}

	newUser, err := h.service.Register(user.Name, user.Email, user.Password, user.Age)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return SendTokenResponse(c, *newUser, fiber.StatusCreated)

}

// Note: This function is not part of the original code snippet
func SendTokenResponse(c *fiber.Ctx, user entity.User, statusCode int) error {
	token, err := authSer.GenerateToken(strconv.FormatInt(int64(user.ID), 10))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Fail to generate token",
		})
	}
	cookieExpire := time.Now().Add(24 * time.Hour) // Default 1 day, adjust as needed

	// Check for custom expiration from environment variables
	if os.Getenv("JWT_COOKIE_EXPIRE") != "" {
		if duration, err := time.ParseDuration(os.Getenv("JWT_COOKIE_EXPIRE")); err == nil {
			cookieExpire = time.Now().Add(duration)
		}
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  cookieExpire,
		HTTPOnly: true,
	})
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}
