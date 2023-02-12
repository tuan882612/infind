package routes

import (
	"auth_ms/api/domain/base"

	"github.com/gin-gonic/gin"
)

func BaseRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/api/v1", base.Base)
	router.NoRoute(base.None)

	return router
}