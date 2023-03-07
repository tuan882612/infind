package test

import (
	"userms/internal/api/v1"
	"userms/pkg/config"

	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := v1.InitService("localhost", 2000)
	server.LoadConfig()
	config.SetEnvVariables(EnvPath)

	return server.Engine
}