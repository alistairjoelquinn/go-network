package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/alistairjoelquinn/go-network/database"
	"github.com/alistairjoelquinn/go-network/model"
	"github.com/alistairjoelquinn/go-network/util"
	"github.com/gofiber/fiber/v2"
	"github.com/pascaldekloe/jwt"
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

	log.Println("userid in check status", userId)

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
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(id)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "localhost:3000"
	claims.Audiences = []string{"localhost:3000"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(tokenSecret.value))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = string(jwtBytes)
	c.Cookie(cookie)

	return c.JSON(fiber.Map{"success": "true"})
}

func LogUserIn(c *fiber.Ctx) error {
	l := new(model.NewUser)

	if err := c.BodyParser(l); err != nil {
		return err
	}

	loginVals, err := database.DBModel.GetUserPasswordFromEmail(l.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}

	log.Println(loginVals.HashedPassword, l.Password)

	err = bcrypt.CompareHashAndPassword([]byte(loginVals.HashedPassword), []byte(l.Password))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}

	err = util.SetTokenAsCookie(c, loginVals.ID)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": "false",
		})
	}

	return c.JSON(fiber.Map{"success": "true"})
}

func LogUserOut(c *fiber.Ctx) error {
	log.Println("LOGOUT")
	c.ClearCookie("token")

	return c.JSON(fiber.Map{"logout": "true"})
}

func CheckEmailForReset(c *fiber.Ctx) error {
	log.Println("check email for reset")
	return nil
}

func VerifyAndResetUsersPassword(c *fiber.Ctx) error {
	log.Println("verify and reset users password")
	return nil
}
