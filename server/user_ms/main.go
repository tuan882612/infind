package main

import (
	"userms/api/v1/routes"
	"userms/api/v1/validators"
	"userms/assets/config"
	// "github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode) // Setting to turnoff debug mode
	config.SetEnvVariables()
	validators.ValidateDynamo()

	router := routes.Init_Service()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	router.Run("0.0.0.0:1000")
}