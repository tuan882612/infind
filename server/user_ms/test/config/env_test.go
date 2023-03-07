package test

import (
	"testing"
	"userms/pkg/config"
	"userms/pkg/validators"
	"userms/test/utility"

	"github.com/stretchr/testify/assert"
)

func Test_LoadEnv(t *testing.T) {
	config.SetEnvVariables(test.EnvPath)
	assert.Equal(t, true, validators.ValidateLoadEnv())
}