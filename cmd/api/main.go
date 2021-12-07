package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})

	app.Static("/", "./public")

	auth := app.Group("/auth")
	auth.Get("/user/id.json", checkUserStatus)

	user := app.Group(("/user"))
	user.Get("/get-data", getUserData)

	log.Fatal(app.Listen(":3001"))
}

func routeErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(404).SendString("Error in route")
}
