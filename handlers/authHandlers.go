package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func CheckUserStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"userId": 45,
	})
}

func GetUserData(c *fiber.Ctx) error {
	log.Println("getting user data")
	return nil
}
