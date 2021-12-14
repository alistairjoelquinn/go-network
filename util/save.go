package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveImageFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "error processing image file", "data": nil})

	}

	n := fmt.Sprintf("%s.%s", uuid.New(), strings.Split(file.Filename, ".")[1])

	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", n))
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "error saving file to disk", "data": nil})
	}

	localTempUrl := fmt.Sprintf("http://localhost:4000/uploads/%s", n)

	log.Println("LOCAL", localTempUrl)

	data := map[string]interface{}{
		"imageName": n,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Upload successful", "data": data})
}
