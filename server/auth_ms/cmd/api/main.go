package main

import (
	"authms/api/server"
	"authms/utility/validators"
	"runtime"
)

func main() {
	router := api.InitService()

	validators.ValidateRedis()
	
	address := "0.0.0.0"
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		address = "localhost"
	}

	router.Run(address+":1003")
}