package routes

import (
	"userms/api/v1/controllers"
	"github.com/gin-gonic/gin"
)

var BaseUrl string = "/api/v1"

func Init_Service() *gin.Engine {
	router := gin.Default()

	router.NoRoute(controllers.None())
	router.GET(BaseUrl, controllers.Base())

	user := router.Group(BaseUrl + "/user") 
	{
		user.GET("/find", controllers.GetUser())
		user.POST("/insert", controllers.CreateUser())
		user.PATCH("/update", controllers.UpdateUser())
		user.DELETE("/delete", controllers.DeleteUser())

		user.GET("/login", controllers.Login())
	}

	log := router.Group(BaseUrl + "/log")
	{
		log.GET("", )
	}

	return router
}