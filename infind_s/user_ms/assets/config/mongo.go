package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Connect_Mongo() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal("Error loading .env file")
    }
}