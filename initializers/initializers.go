package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
