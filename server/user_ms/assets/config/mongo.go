package config

import (
	"log"

	"github.com/joho/godotenv"
)

func ConnectMongo() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal("Error loading .env file")
    }
}