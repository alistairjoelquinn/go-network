package handler

import (
	"log"
	"time"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/model"
	"github.com/alistairjoelquinn/go-network/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type jwtBuild struct {
	value string
}

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

	id, err := database.DBModel.AddNewUser(n.First, n.Last, n.Email, string(hashedPassword))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["firstname"] = n.First
	claims["lastname"] = n.Last
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenSecret := jwtBuild{
		value: util.Env("JWT_SECRET"),
	}

	t, err := token.SignedString([]byte(tokenSecret.value))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	c.Cookie(cookie)

	return c.JSON(fiber.Map{"success": "true"})
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
