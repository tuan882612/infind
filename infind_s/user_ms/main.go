package main

import (
	"userms/api/v1/routes"
	"userms/assets/config"

	// "github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode) // Setting to turnoff debug mode

	config.Connect_Mongo()
	
	router := routes.Init_Service()
	
	router.SetTrustedProxies([]string{"192.168.1.2"})

	router.Run("localhost:1000")
}