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
	PicturesRoutes.Post("/", PictureHandler.UploadPics)
	PicturesRoutes.Get("/:id", PictureHandler.GetPicsByUserId)
}
