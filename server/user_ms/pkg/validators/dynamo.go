package validators

import (
	"log"
	"userms/pkg/config"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ValidateDynamo() bool {
	db := config.ConnectDynamodb()
	_, err := db.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		log.Fatalln("Unable to connect to AWS.")
	}
	println("Connected to AWS.")
	
	return true
}