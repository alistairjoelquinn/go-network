package handler

import (
	"log"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/models"
	"github.com/gofiber/fiber/v2"
)

func CheckUserStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		// "userId": 45,
	})
}

func CreateNewUser(c *fiber.Ctx) error {
	n := new(models.NewUser)

	if err := c.BodyParser(n); err != nil {
		return err
	}

	err := database.DBModel.AddNewUser(n)
	if err != nil {
		log.Println("Failed")
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}
	log.Println("succeedededed")

	return c.JSON(fiber.Map{
		"success": "true",
	})
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
