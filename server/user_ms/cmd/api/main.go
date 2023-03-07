package main

import (
	"user_ms/internal/api/v1"
	"user_ms/pkg/validators"
	"user_ms/pkg/config"

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
