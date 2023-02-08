package routes

import (
	"userms/api/v1/domain/login"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(router *gin.Engine) *gin.Engine {
	handler := &login.LoginController{
		Repo: login.NewRepo(),
	}
	
	login := router.Group("/api/v1")
	{
		login.GET("/login", handler.Login)
		login.POST("/register", handler.Register)
	}

	return router
}