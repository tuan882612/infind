package validators

import (
	"log"
	"userms/assets/config"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ValidateDynamo() {
	db := config.ConnectDynamodb()
	_, err := db.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		log.Fatalln("Could not connect to AWS.")
	}
}