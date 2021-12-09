package main

import (
	"log"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	_ "github.com/lib/pq"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})
	cryptoKey := encryptcookie.GenerateKey()

	app.Use(func(c *fiber.Ctx) error {
		log.Println(c.Cookies("userId", "not working"))
		return c.Next()
	})
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: cryptoKey,
	}))

	app.Static("/", "./public")

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func routeErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(404).SendString("Error in route")
}
