package routes

import (
	"userms/api/v1/controllers"
	"github.com/gin-gonic/gin"
)

func Init_Service() *gin.Engine {
	router := gin.Default()

	router.NoRoute(controllers.None())

	api := router.Group("/api/v1") 
	{
		api.GET("/", controllers.Base())
	}

	return router
}