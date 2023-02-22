package routes

import (
	"authms/api/domain/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) *gin.Engine {
	handler := &auth.AuthController{
		AuthRepo: auth.NewAuthRepo(),
	}
	
	api := router.Group("/api/v1/auth")
	{
		api.GET("/verify")
		api.POST("/create", handler.CreateToken)
	}

	return router
}