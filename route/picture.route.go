package route

import (
	"fmt"

	"github.com/Dpyde/Omchu/middleware"
	"github.com/Dpyde/Omchu/picture"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupPictureRoutes(app *fiber.App, db *gorm.DB) {
	PictureRepo := picture.NewGormPictureRepository(db)
	PictureService := picture.NewPictureService(PictureRepo)
	PictureHandler := picture.NewHttpPictureHandler(PictureService)
	fmt.Println("Kuy")

	PicturesRoutes := app.Group("/picture")
	PicturesRoutes.Use(middleware.Middleware)
	// PicturesRoutes.Post("/", PictureHandler.UploadPictures)
	// PicturesRoutes.Get("/:id", PictureHandler.GetPicture)
	PicturesRoutes.Post("/:id", func(c *fiber.Ctx) error {
		fmt.Println("UploadPictures route hit!") // Log to confirm if the route is being hit
		return PictureHandler.UploadPicture(c)
	})
}
