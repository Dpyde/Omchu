package authHndl

import (
	"os"
	"strconv"
	"time"

	"github.com/Dpyde/Omchu/internal/entity"
	auth "github.com/Dpyde/Omchu/internal/service/auth"
	"github.com/gofiber/fiber/v2"
)

func SendTokenResponse(c *fiber.Ctx, user entity.User, statusCode int) error {
	token, err := auth.GenerateToken(strconv.FormatInt(int64(user.ID), 10))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"msg":     "Fail to generate token",
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
