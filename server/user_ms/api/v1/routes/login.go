package routes

import (
	"userms/api/v1/controllers"
	"userms/api/v1/repository"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.Engine) *gin.Engine {
	handler := &controllers.LoginController{
		Repo: repository.NewRepo(),
	}
	
	login := router.Group("/api/v1")
	{
		login.GET("/login", handler.Login)
		login.POST("/register", handler.Register)
	}

	return router
}