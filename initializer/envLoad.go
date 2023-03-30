package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if godotenv.Load() != nil {
		panic("Failed to load env file")
	}
	log.Println("env file loaded successfully!")
}
