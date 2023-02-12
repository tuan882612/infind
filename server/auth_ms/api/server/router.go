package api

import (
	"auth_ms/api/server/routes"
	"auth_ms/assets/config"

	"github.com/gin-gonic/gin"
)

func InitService() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	
	config.SetEnvVariables()
	config.Middleware(router)

	routes.BaseRoutes(router)
	routes.AuthRoutes(router)

	return router
}