package validators

import (
	"log"
	"os"
)

func ValidateLoadEnv() bool {
	if os.Getenv("test") != "foo" {
		log.Fatal("Unable to load environment variables.")
	}
	println("Loaded environment variables.")
	
	return true
}