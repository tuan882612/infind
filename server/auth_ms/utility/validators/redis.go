package validators

import (
	"auth_ms/assets/config"
	"log"
)

func ValidateRedis() {
	client := config.ConnectRedis()

	ctx := client.Context()
	pong, err := client.Ping(ctx).Result()

	if pong == "" || err != nil {
		log.Fatalln("Unable to connect to redis.")
	} else {
		println("Connected to redis.")
	}
}