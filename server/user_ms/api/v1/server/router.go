package v1

import (
	"userms/api/v1/server/routes"
	"userms/assets/config"

	"github.com/gin-gonic/gin"
)

func InitService() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	config.SetEnvVariables()
	config.MiddleWare(router)
	
	routes.Default(router)
	routes.UserRoutes(router)
	routes.LoginRoutes(router)

	return router
}