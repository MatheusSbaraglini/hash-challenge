package env_test

import (
	"log"
	"os"
	"testing"

	"github.com/matheussbaraglini/hash-challenge/pkg/env"
	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	t.Run("should return value from environment var", func(t *testing.T) {
		envVar := "VAR_TEST"
		err := os.Setenv(envVar, "value_from_test")
		assert.NoError(t, err)

		defer os.Unsetenv(envVar)
		value := env.GetString(envVar)
		assert.Equal(t, "value_from_test", value)
	})

	t.Run("should return default value", func(t *testing.T) {
		envVar := "VAR_TEST"

		defer os.Unsetenv(envVar)
		value := env.GetString(envVar, "default_value")
		assert.Equal(t, "default_value", value)
	})
}

func TestCheckRequired(t *testing.T) {
	testLog := log.New(os.Stderr, "", log.LstdFlags)

	t.Run("should validate all environment variables required", func(t *testing.T) {
		envVar := "VAR_TEST"
		err := os.Setenv(envVar, "value_from_test")
		assert.NoError(t, err)

		defer os.Unsetenv(envVar)
		err = env.CheckRequired(*testLog, envVar)
		assert.NoError(t, err)
	})

	t.Run("should validate all environment variables required and return error", func(t *testing.T) {
		err := os.Setenv("VAR_TEST2", "value_from_test")
		assert.NoError(t, err)
		defer os.Unsetenv("VAR_TEST2")

		err = env.CheckRequired(*testLog, "VAR_TEST", "VAR_TEST1", "VAR_TEST2")
		assert.Error(t, err)
		assert.EqualError(t, err, "environment variables are required: [\"VAR_TEST\" \"VAR_TEST1\"]")
	})
}
