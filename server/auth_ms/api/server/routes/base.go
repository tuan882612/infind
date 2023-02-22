package routes

import (
	"authms/api/domain/base"

	"github.com/gin-gonic/gin"
)

func BaseRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/api/v1", base.Default)
	router.NoRoute(base.None)

	return router
}