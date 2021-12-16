package handler

import (
	"log"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/model"
	"github.com/alistairjoelquinn/go-network/util"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type jwtBuild struct {
	value string
}

var tokenSecret = jwtBuild{
	value: util.Env("JWT_SECRET"),
}

func CheckUserStatus(c *fiber.Ctx) error {
	userId, err := util.GetIdFromToken(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"userId": "",
		})
	}

	return c.JSON(fiber.Map{
		"userId": userId,
	})
}

func CreateNewUser(c *fiber.Ctx) error {
	n := new(model.NewUser)

	if err := c.BodyParser(n); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	id, err := database.DBModel.AddNewUser(n.First, n.Last, n.Email, string(hashedPassword))
	if err != nil {
		return c.SendStatus(500)
	}

	err = util.SetTokenAsCookie(c, id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{"success": "true"})
}

func LogUserIn(c *fiber.Ctx) error {
	l := new(model.NewUser)

	if err := c.BodyParser(l); err != nil {
		return err
	}

	loginVals, err := database.DBModel.GetUserPasswordFromEmail(l.Email)
	if err != nil {
		return c.SendStatus(401)
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginVals.HashedPassword), []byte(l.Password))
	if err != nil {
		return c.SendStatus(401)
	}

	err = util.SetTokenAsCookie(c, loginVals.ID)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{"success": "true"})
}

func LogUserOut(c *fiber.Ctx) error {
	c.ClearCookie("token")

	return c.JSON(fiber.Map{"logout": "true"})
}

func CheckEmailForReset(c *fiber.Ctx) error {
	l := new(model.CheckEmail)

	if err := c.BodyParser(l); err != nil {
		return c.SendStatus(500)
	}

	err := database.DBModel.CheckEmailForReset(l.Email)
	if err != nil {
		return c.SendStatus(500)
	}

	code, err := util.GenerateRandomString(6)
	if err != nil {
		return c.SendStatus(500)
	}

	log.Println("random string to use a code", code)

	return nil
}

func VerifyAndResetUsersPassword(c *fiber.Ctx) error {
	log.Println("verify and reset users password")
	return nil
}
