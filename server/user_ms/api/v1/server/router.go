package v1

import (
	"userms/api/v1/server/routes"
	"userms/assets/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitService() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	config.SetEnvVariables()

	router := gin.Default()
	router.Use(cors.New(config.MiddleWare()))
	
	routes.Default(router)
	routes.UserRoutes(router)
	routes.LoginRoutes(router)

	return router
}