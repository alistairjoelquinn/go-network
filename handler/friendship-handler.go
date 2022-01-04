package handler

import (
	"log"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/util"
	"github.com/gofiber/fiber/v2"
)

func GetInitialFrienshipStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.SendStatus(401)
	}

	friendshipStatus, err := database.DBModel.FriendshipStatus(userId, id)
	if err != nil {
		return c.JSON(nil)
	}
	return c.JSON(friendshipStatus)
}

func AddFriend(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.SendStatus(401)
	}

	addFriendNewStatus, err := database.DBModel.AddFriendQuery(userId, id)
	if err != nil {
		log.Println(err, "ERROR")
		return c.SendStatus(500)
	}
	return c.JSON(addFriendNewStatus)
}

func AcceptFriend(c *fiber.Ctx) error {
	// id := c.Params("id")

	return nil
}

func EndFriendship(c *fiber.Ctx) error {
	// id := c.Params("id")

	return nil
}

func GetFriendsList(c *fiber.Ctx) error {
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.SendStatus(401)
	}

	users, err := database.DBModel.GetRequestsFriends(userId)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(users)
}
