package main

import (
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
	router.Run("0.0.0.0:1000")
}