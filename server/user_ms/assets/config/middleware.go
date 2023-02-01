package config

import (
	"github.com/gin-contrib/cors"
)

func MiddleWare() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AddAllowHeaders(
		"Access-Control-Allow-Headers", 
		"access-control-allow-origin, access-control-allow-headers", 
		"Content-Type", "X-XSRF-TOKEN", 
		"Accept", "Origin", "X-Requested-With", "Authorization")
	config.AddAllowMethods("*")
	return config
}