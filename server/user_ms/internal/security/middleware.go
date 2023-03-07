package security

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MiddleWare(router *gin.Engine) *gin.Engine {
	origins := []string{
		"http://localhost:3000",
		"http://localhost:1001",
		"http://localhost:1002",
		"http://localhost:1003",
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = origins
	config.AllowCredentials = true
	config.AddAllowHeaders(
		"Access-Control-Allow-Headers",
		"access-control-allow-origin, access-control-allow-headers",
		"Content-Type", "X-XSRF-TOKEN",
		"Accept", "Origin", "X-Requested-With", "Authorization")
	config.AddAllowMethods("*")
	router.Use(cors.New(config))
	
	return router
}