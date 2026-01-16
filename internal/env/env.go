package env

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func GetString(key, fallback string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return fallback
}
