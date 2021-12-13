package util

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pascaldekloe/jwt"
)

type jwtBuild struct {
	value string
}

var tokenSecret = jwtBuild{
	value: Env("JWT_SECRET"),
}

func GetIdFromToken(c *fiber.Ctx) (string, error) {
	tokenErr := errors.New("error checking identity")

	token := c.Cookies("token", "")
	if token == "" {
		return "", tokenErr
	}

	claims, err := jwt.HMACCheck([]byte(token), []byte(tokenSecret.value))
	if err != nil || !claims.Valid(time.Now()) || !claims.AcceptAudience("localhost:3000") || claims.Issuer != "localhost:3000" {
		c.ClearCookie("token")
		return "", tokenErr
	}

	_, err = strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return "", tokenErr
	}

	return claims.Subject, nil
}

func SetTokenAsCookie(c *fiber.Ctx, id string) error {
	setTokenError := errors.New("error setting token")

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(id)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "localhost:3000"
	claims.Audiences = []string{"localhost:3000"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(tokenSecret.value))
	if err != nil {
		return setTokenError
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = string(jwtBytes)
	c.Cookie(cookie)

	return nil
}
