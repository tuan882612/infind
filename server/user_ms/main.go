package main

import (
	"runtime"
	"userms/api/v1"
	"userms/api/validators"
	"userms/assets/config"
	// "github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode) // Setting to turnoff debug mode
	config.SetEnvVariables()
	validators.ValidateDynamo()

	router := v1.InitService()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	address := "0.0.0.0"

	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		address = "localhost"
	}

	router.Run(address + ":1000")
}
