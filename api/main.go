package main

import (
	"log"

	"github.com/alistairjoelquinn/go-network/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: routeErrorHandler,
	})

	app.Static("/", "./public")

	auth := app.Group("/auth")
	auth.Get("/user/id.json", handlers.CheckUserStatus)
	auth.Post("/register", handlers.CreateNewUser)
	auth.Post("/login", handlers.LogUserIn)
	auth.Get("/logout", handlers.LogUserOut)
	auth.Get("/password-reset/email-check", handlers.CheckEmailForReset)
	auth.Get("/password-reset/verify-code", handlers.VerifyAndResetUsersPassword)

	user := app.Group(("/user"))
	user.Get("/get-data", handlers.GetUserData)
	user.Get("/upload", handlers.UploaderUserImage)
	user.Get("/set-bio", handlers.SetUserBio)
	user.Get("/recent-users", handlers.GetRecentUsers)
	user.Get("/user-search/:q", handlers.SearchForUsers)
	user.Get("/other-user/:id", handlers.GetOtherUser)

	// friendship := app.Group(("/friendship"))

	log.Fatal(app.Listen(":3000"))
}

func routeErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(404).SendString("Error in route")
}
