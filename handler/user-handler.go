package handler

import (
	"log"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/model"
	"github.com/alistairjoelquinn/go-network/util"
	"github.com/gofiber/fiber/v2"
)

func GetUserData(c *fiber.Ctx) error {
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.SendStatus(500)
	}

	user, err := database.DBModel.GetUserData(userId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(user)
}

func UploaderUserImage(c *fiber.Ctx) error {
	image, err := util.SaveImageFile(c)
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "error saving file to disk", "data": nil})
	}

	imageUrl, err := util.UploadImage(image)
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "error uploading image to cloud", "data": nil})
	}

	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "error getting user ID", "data": nil})

	}

	err = database.DBModel.AddNewUserImage(userId, imageUrl)
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "error saving image to database", "data": nil})
	}

	return c.JSON(fiber.Map{
		"image": imageUrl,
	})
}

func SetUserBio(c *fiber.Ctx) error {
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.SendStatus(401)
	}

	b := new(model.UpdatedBio)
	if err := c.BodyParser(b); err != nil {
		return c.SendStatus(404)
	}

	err = database.DBModel.UpdateUserBio(userId, b.Bio)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(b)
}

func GetRecentUsers(c *fiber.Ctx) error {
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.SendStatus(401)
	}

	users, err := database.DBModel.RecentUserSearch(userId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(users)
}

func SearchForUsers(c *fiber.Ctx) error {
	q := c.Params("q")

	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.SendStatus(401)
	}
	log.Println("q, userId", q, userId)

	users, err := database.DBModel.UserSearch(q, userId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(users)
}

func GetOtherUser(c *fiber.Ctx) error {
	log.Println("get other user")
	return nil
}
