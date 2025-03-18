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

func (h *HttpPictureHandler) UploadPics(c *fiber.Ctx) error {
	// Parse user ID from form-data
	idStr := c.FormValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid user ID",
		})
	}

	// Get multiple files from form-data
	form, err := c.MultipartForm()
	if err != nil || form.File["picture"] == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "no files uploaded",
		})
	}

	files := form.File["picture"]

	// Upload the pictures to Cloudflare R2
	pictures, err := h.service.UploadPicsToR2(files, os.Getenv("BUCKET_NAME"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to upload files",
		})
	}

	// Assign user ID to each picture and save to DB
	for i := range pictures {
		pictures[i].UserID = uint(id)
	}

	// Optionally save the pictures to the database here (commented out for now)
	if err := h.service.SavePicturesSer(uint(id), pictures); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to save pictures",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"pictures": pictures,
	})
}

func (h *HttpPictureHandler) GetPicsByUserId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	pictures, err := h.service.GetPicsByUserId(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "failed to get Pictures from database"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"pictures": pictures,
	})
}
