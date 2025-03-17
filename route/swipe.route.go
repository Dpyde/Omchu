package route

import (
	swipeGormRep "github.com/Dpyde/Omchu/adapter/gorm/swipe"
	swipeHndl "github.com/Dpyde/Omchu/adapter/http/swipe"
	swipeSer "github.com/Dpyde/Omchu/internal/service/swipe"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupSwipeRoutes(app *fiber.App, db *gorm.DB) {
	swipeRepo := swipeGormRep.NewGormSwipeRepository(db)
	swipeService := swipeSer.NewSwipeService(swipeRepo)
	swipeHandler := swipeHndl.NewHttpSwipeHandler(swipeService)

	swipesRoutes := app.Group("/swipe")
	swipesRoutes.Post("/", swipeHandler.SwipeCheck)
}
