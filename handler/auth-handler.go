package handler

import (
	"log"
	"time"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"honnef.co/go/tools/config"
)

func CheckUserStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		// "userId": 45,
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

	err = database.DBModel.AddNewUser(n.First, n.Last, n.Email, string(hashedPassword))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
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
