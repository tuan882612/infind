package main

import (
	"userms/internal/api/v1"
	"userms/pkg/validators"
	"userms/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	config.SetEnvVariables(".env")
	validators.ValidateLoadEnv()
	validators.ValidateDynamo()

	server := v1.InitService("localhost",2000)
	server.LoadConfig()
	server.Run()
}
