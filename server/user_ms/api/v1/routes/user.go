package routes

import (
	"userms/api/v1/controllers"
	"userms/api/v1/repository"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) *gin.Engine {
	handler := &controllers.UserController{
		Repo: repository.NewRepo(),
	}
	
	user := router.Group("/api/v1/user") 
	{
		user.GET("/find", handler.Find)
		user.PUT("/update", handler.Update)
		user.DELETE("/delete", handler.Delete)
	}

	return router
}