package dotenv

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Load auto-load env variables from .env file in the root of project
func Load(filenames ...string) error {
	err := godotenv.Load(filenames...)

	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	return nil
}

func GetString(varName string, defaultValue string) string {
	if envVal, ok := os.LookupEnv(varName); ok {
		return envVal
	}

	return defaultValue
}

func GetInt(varName string, defaultValue int) int {
	envVal, ok := os.LookupEnv(varName)

	if !ok {
		return defaultValue
	}

	if num, err := strconv.Atoi(envVal); err == nil {
		return num
	}

	return defaultValue
}
