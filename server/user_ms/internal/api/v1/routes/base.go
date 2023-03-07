package routes

import (
	h "user_ms/internal/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func Default(router *gin.RouterGroup) {
	router.GET("", h.Default)
	router.GET("/health", h.Health)
}