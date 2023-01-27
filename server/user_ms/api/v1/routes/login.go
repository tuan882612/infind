package routes

import (
	"userms/api/v1/controllers"
	"userms/api/v1/database"
	// "userms/assets/config"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.Engine) *gin.Engine {
	handler := &controllers.LoginController{
		Repo: database.NewRepo(),
	}
	
	login := router.Group("/api/v1")
	{
		login.GET("/login", handler.Login)
		login.POST("/register", handler.Register)
	}

	return router
}