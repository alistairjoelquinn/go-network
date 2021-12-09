package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env")
	}
	return os.Getenv(key)
}
