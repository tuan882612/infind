package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Connect_Dyanmo() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal("Error loading .env file")
    }
}