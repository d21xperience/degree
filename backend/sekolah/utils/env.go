package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnvFiles() {
	envFiles := []string{".env.local", ".env"}

	for _, envFile := range envFiles {
		err := godotenv.Load(envFile)
		if err == nil {
			log.Printf("Loaded environment from: %s", envFile)
			return
		}
	}

	log.Println("No .env files found, using OS environment variables")
}

// Helper functions
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
