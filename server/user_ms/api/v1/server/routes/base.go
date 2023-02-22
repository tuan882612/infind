package routes

import (
	"userms/api/v1/domain/base"

	"github.com/gin-gonic/gin"
)

func Default(router *gin.Engine) *gin.Engine {
	router.GET("/api/v1", base.Default)
	router.NoRoute(base.None)

	return router
}