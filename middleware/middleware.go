package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Middleware validates JWT, extracts userId, and stores it in Fiber's context
func Middleware(c *fiber.Ctx) error {
	// Retrieve the token from the cookie
	fmt.Println("Middleware executed")
	tokenString := c.Cookies("token")

	// If there's no token in the cookie, return unauthorized
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Token mueng mai mee wa",
		})
	}

	// Retrieve JWT secret key
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "JWT secret not configured",
		})
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Token mueng mai mee wa",
		})
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid token claims",
		})
	}

	// Extract user ID from the claims
	userID, ok := claims["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID not found in token",
		})
	}

	// Store the userId in Fiber's context
	c.Locals("UserId", userID)

	// If everything is fine, allow the request to continue
	return c.Next()
}
