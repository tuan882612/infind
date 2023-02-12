package main

import (
	api "auth_ms/api/server"
	"auth_ms/utility/validators"
	"runtime"
)

func main() {
	validators.ValidateRedis()

	router := api.InitService()

	address := "0.0.0.0"
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		address = "localhost"
	}

	router.Run(address+":1003")
}