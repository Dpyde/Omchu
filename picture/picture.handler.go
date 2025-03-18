package picture

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HttpPictureHandler struct {
	service PictureService
}

func NewHttpPictureHandler(service PictureService) *HttpPictureHandler {
	return &HttpPictureHandler{service: service}
}

// func (h *HttpPictureHandler) UploadPictures(c *fiber.Ctx) error {
// 	// Parse user ID from form-data
// 	idStr := c.FormValue("id")
// 	id, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid user ID"})
// 	}

// 	// Get multiple files
// 	form, err := c.MultipartForm()
// 	if err != nil || form.File["picture"] == nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "no files uploaded"})
// 	}
// 	files := form.File["picture"]

// 	// Upload files to R2
// 	pictures, err := h.service.UploadPicsToR2(files, os.Getenv("BUCKET_NAME"))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "failed to upload files"})
// 	}

// 	// Assign user ID and save to DB
// 	for i := range pictures {
// 		pictures[i].UserID = uint(id)
// 	}
// 	if err := h.service.SavePicturesSer(uint(id), pictures); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "failed to save pictures"})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "pictures": pictures})
// }

func (h *HttpPictureHandler) UploadPicture(c *fiber.Ctx) error {
	// Parse user ID from form-data
	idStr := c.FormValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid user ID",
		})
	}

	// Get the single file from form-data
	file, err := c.FormFile("picture")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "no file uploaded",
		})
	}

	// Upload the picture to Cloudflare R2
	picture, err := h.service.UploadPicToR2(file, os.Getenv("BUCKET_NAME"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to upload file",
		})
	}

	picture.UserID = uint(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"picture": picture,
	})
}

func (h *HttpPictureHandler) GetPicturesByUserId(c *fiber.Ctx) error {
	var id uint
	if err := c.BodyParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request"})
	}

	pictures, err := h.service.GetPicturesByUserId(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "failed to get Pictures from database"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"pictures": pictures,
	})
}
