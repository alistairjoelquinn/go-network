package handler

import (
	"log"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/util"
	"github.com/gofiber/fiber/v2"
)

func GetUserData(c *fiber.Ctx) error {
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.JSON(fiber.Map{"success": "false"})
	}

	user, err := database.DBModel.GetUserData(userId)
	if err != nil {
		return c.JSON(fiber.Map{"success": "false"})
	}
	log.Println("user", user)

	return c.JSON(fiber.Map{
		"success": "true",
		"data":    user,
	})
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
