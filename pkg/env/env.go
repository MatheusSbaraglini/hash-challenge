package env

import (
	"fmt"
	"log"
	"os"
)

func GetString(envVar string, defaultValue ...string) string {
	value := os.Getenv(envVar)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}

func CheckRequired(log log.Logger, envVars ...string) error {
	var requiredVars []string

	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			requiredVars = append(requiredVars, envVar)
			continue
		}

		log.Printf("Environment variable %s is ok\n", envVar)
	}

	if len(requiredVars) > 0 {
		return fmt.Errorf("environment variables are required: %+q", requiredVars)
	}

	return nil
}
