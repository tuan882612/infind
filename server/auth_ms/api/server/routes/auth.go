package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.Engine) *gin.Engine {
	api := router.Group("/auth")
	{
		api.GET("")
		api.POST("")
	}

	return router
}