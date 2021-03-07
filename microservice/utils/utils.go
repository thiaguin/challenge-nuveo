package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVariable func
func GetEnvVariable(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
