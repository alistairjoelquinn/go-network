package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/alistairjoelquinn/go-network/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	_ "github.com/lib/pq"
)

var cryptoKey = encryptcookie.GenerateKey()

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
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

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/sn-typescript?sslmode=disable")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
