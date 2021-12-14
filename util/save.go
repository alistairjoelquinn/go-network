package util

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveImageFile(c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return "", err
	}

	n := fmt.Sprintf("%s.%s", uuid.New(), strings.Split(file.Filename, ".")[1])

	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", n))
	if err != nil {
		return "", err
	}

	return n, nil
}
