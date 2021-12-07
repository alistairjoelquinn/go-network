package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})
	store := session.New()

	app.Use("/", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		sess.Set("userId", 56)
		userId := sess.Get("userId")

		log.Println("session value:", userId)
		return c.Next()
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
