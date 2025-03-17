package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupSwipeRoutes(app *fiber.App, db *gorm.DB) {
	swipeRepo := swipe.NewGormSwipeRepository(db)
	swipeService := swipe.NewSwipeService(swipeRepo)
	swipeHandler := swipe.NewHttpSwipeHandler(swipeService)

	swipesRoutes := app.Group("/swipe")
	swipesRoutes.Post("/", swipeHandler.SwipeCheck)
}
