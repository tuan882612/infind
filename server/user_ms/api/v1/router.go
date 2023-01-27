package v1

import (
	"userms/api/v1/controllers"
	"userms/api/v1/routes"

	"github.com/gin-gonic/gin"
)

var BaseUrl string = "/api/v1"

func InitService() *gin.Engine {
	router := gin.Default()
	router.NoRoute(controllers.None)
	router.GET(BaseUrl, controllers.Base)

	routes.UserRoutes(router)
	routes.LoginRoutes(router)

	return router
}