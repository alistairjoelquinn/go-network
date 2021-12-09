package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(c *fiber.Ctx) error {
	log.Println("getting user data")
	return nil
}

func UploaderUserImage(c *fiber.Ctx) error {
	log.Println("upload user image")
	return nil
}

func SetUserBio(c *fiber.Ctx) error {
	log.Println("set user bio")
	return nil
}

func GetRecentUsers(c *fiber.Ctx) error {
	log.Println("get recent users")
	return nil
}

func SearchForUsers(c *fiber.Ctx) error {
	log.Println("searach for users")
	return nil
}

func GetOtherUser(c *fiber.Ctx) error {
	log.Println("get other user")
	return nil
}
