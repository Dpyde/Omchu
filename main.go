package main

import (
	"fmt"
	"log"

	"github.com/Dpyde/Omchu/database"
	"github.com/Dpyde/Omchu/picture"
	"github.com/Dpyde/Omchu/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	app := fiber.New()
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	if err1 := godotenv.Load(); err1 != nil {
		log.Fatal(err1)
	}

	picture.InitR2()

	fmt.Println("Database connected")
	// Configure your PostgreSQL database details here

	route.SetupPictureRoutes(app, db)
	route.SetupChatRoutes(app, db)
	route.SetupUserRoutes(app, db)
	route.SetupAuthRoutes(app, db)
	route.SetupSwipeRoutes(app, db)
	// app.Post("/test", image.TestSend)
	// Endpoint to upload image
	// app.Post("/upload", func(c *fiber.Ctx) error {
	// 	fileHeader, err := c.FormFile("file")
	// 	if err != nil {
	// 		return c.Status(fiber.StatusBadRequest).SendString("Failed to get file")
	// 	}

	// 	fileURL, fileKey, err := cloudflare.UploadFileToR2(fileHeader, os.Getenv("BUCKET_NAME"))
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to upload file")
	// 	}

	// 	// Save picture details to the database
	// 	picture := entity.Picture{
	// 		UserID: 1, // Replace with the actual user ID
	// 		Url:    fileURL,
	// 		Key:    fileKey,
	// 	}
	// 	if err := db.Create(&picture).Error; err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save picture details")
	// 	}

	// 	return c.SendString("File uploaded and saved successfully")
	// })

	// Start the server
	app.Listen(":8000")
}
