package main

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()
	app.Post("/api/upload", uploadFiles)
	app.Listen(":8080")
}

func uploadFiles(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Bad Request",
		})
	}

	files := form.File["images"]
	fileNameSlice := []string{}

	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)
		fileNameSlice = append(fileNameSlice, newFileName)

		if err := c.SaveFile(file, "./uploads/"+newFileName); err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Error While Save Files",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fileName": fileNameSlice,
	})
}
