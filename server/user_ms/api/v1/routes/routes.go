package routes

import (
	"userms/api/v1/controllers"
	"github.com/gin-gonic/gin"
)

var ApiVersion string = "/api/v1/"

func Init_Service() *gin.Engine {
	router := gin.Default()

	router.NoRoute(controllers.None())

	user := router.Group(ApiVersion + "user") 
	{
		user.GET("", controllers.Base())
		user.GET("/find", controllers.GetUser())
		user.POST("/insert", controllers.CreateUser())
		user.PATCH("/update", controllers.UpdateUser())
		user.DELETE("/delete", controllers.DeleteUser())
	}

	log := router.Group(ApiVersion + "log")
	{
		log.GET("", controllers.Base())
	}

	return router
}