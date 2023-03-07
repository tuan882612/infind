package test

import (
	"testing"
	"userms/pkg/config"
	"userms/pkg/validators"
	test "userms/test/utility"

	"github.com/stretchr/testify/assert"
)

func Test_DynamodbConnection(t *testing.T) {
	config.SetEnvVariables(test.EnvPath)
	assert.Equal(t, true, validators.ValidateDynamo())
}