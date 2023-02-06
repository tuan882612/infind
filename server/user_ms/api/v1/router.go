package v1

import (
	"userms/api/v1/controllers"
	"userms/api/v1/routes"
	"userms/assets/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var BaseUrl string = "/api/v1"

func InitService() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	config.SetEnvVariables()

	router := gin.Default()
	router.Use(cors.New(config.MiddleWare()))
	
	router.NoRoute(controllers.None)
	router.GET(BaseUrl, controllers.Base)

	routes.UserRoutes(router)
	routes.LoginRoutes(router)

	return router
}