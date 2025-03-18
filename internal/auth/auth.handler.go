package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type HttpAuthHandler struct {
	service AuthService
}

func NewHttpAuthHandler(service AuthService) *HttpAuthHandler {
	return &HttpAuthHandler{service: service}
}

func (h *HttpAuthHandler) Login(c *fiber.Ctx) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}
	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	// return c.JSON(fiber.Map{"token": token})
	return SendTokenResponse(c, user.ID, fiber.StatusOK)
}

func (h *HttpAuthHandler) Register(c *fiber.Ctx) error {
	type Req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Age      uint   `json:"age"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}

	newUser, err := h.service.Register(req.Name, req.Email, req.Password, req.Age)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return SendTokenResponse(c, newUser.ID, fiber.StatusCreated)

}

// Note: This function is not part of the original code snippet
func SendTokenResponse(c *fiber.Ctx, id uint, statusCode int) error {
	token, err := GenerateToken(strconv.FormatInt(int64(id), 10))
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

func RetrieveTokenRequest(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	id, err := TokenToId(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "error": "invalid token"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "id": id})
}
