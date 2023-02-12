package main

import (
	"runtime"
	"userms/api/v1/server"
	"userms/utility/validators"
)

func main() {
	router := v1.InitService()
	
	validators.ValidateDynamo()

	address := "0.0.0.0"
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		address = "localhost"
	}

	router.Run(address+":1000")
}
