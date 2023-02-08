package main

import (
	"runtime"
	"userms/api/v1"
)

func main() {
	router := v1.InitService()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	address := "0.0.0.0"
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		address = "localhost"
	}

	router.Run(address+":1000")
}
