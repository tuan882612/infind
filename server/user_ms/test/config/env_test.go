package test

import (
	"testing"
	"user_ms/pkg/config"
	"user_ms/pkg/validators"
	"user_ms/test/utility"

	"github.com/stretchr/testify/assert"
)

func Test_LoadEnv(t *testing.T) {
	config.SetEnvVariables(test.EnvPath)
	assert.Equal(t, true, validators.ValidateLoadEnv())
}