package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func checkUserStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"userId": 45,
	})
}

func getUserData(c *fiber.Ctx) error {
	log.Println("getting user data")
	return nil
}
