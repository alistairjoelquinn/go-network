package main

import (
	"log"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.ModelInit(db)

	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})

	app.Static("/", "./public")

	router.SetupRoutes(app)

	app.Static("/*", "./public")

	log.Fatal(app.Listen(":3000"))
}

func routeErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(404).SendString("Error in route")
}
