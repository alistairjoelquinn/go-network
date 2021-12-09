package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckUserStatus(c *fiber.Ctx) error {
	userIdCookie := new(fiber.Cookie)
	userIdCookie.Name = "userId"
	userIdCookie.Value = "45"
	userIdCookie.Expires = time.Now().Add(72 * time.Hour)
	c.Cookie(userIdCookie)

	return c.JSON(fiber.Map{
		"userId": 45,
	})
}

func CreateNewUser(c *fiber.Ctx) error {
	log.Println("create new user")
	return nil
}

func LogUserIn(c *fiber.Ctx) error {
	log.Println("log user in")
	return nil
}

func LogUserOut(c *fiber.Ctx) error {
	log.Println("log user out")
	return nil
}

func CheckEmailForReset(c *fiber.Ctx) error {
	log.Println("check email for reset")
	return nil
}

func VerifyAndResetUsersPassword(c *fiber.Ctx) error {
	log.Println("verify and reset users password")
	return nil
}
