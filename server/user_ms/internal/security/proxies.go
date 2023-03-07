package security

import "github.com/gin-gonic/gin"

func SetTrustedProxies(engine *gin.Engine) {
	proxies := []string{
		"192.168.1.2", 
		"localhost",
		"0.0.0.0",
		"127.0.0.1",
	}
	engine.SetTrustedProxies(proxies)
}