package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisHost       string
	RedisPort       int
	RedisDBName     int
	RedisPassword   string
	DBHost          string
	DBPort          int
	DBUser          string
	DBPassword      string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	GRPCHost        string
	GRPCPort        int
	HTTPPort        int
}

func LoadConfig() Config {
	// Load environment variables
	loadEnvFiles()

	return Config{
		// Redis
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getIntEnv("REDIS_PORT", 6379),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDBName:   getIntEnv("REDIS_DB", 0),
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getIntEnv("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "dbsekolah"),
		// Connection Pool
		MaxIdleConns:    getIntEnv("MAX_IDLE_CONNS", 10),
		MaxOpenConns:    getIntEnv("MAX_OPEN_CONNS", 100),
		ConnMaxLifetime: time.Duration(getIntEnv("CONN_MAX_LIFETIME_MINUTES", 30)) * time.Minute,

		// App
		GRPCHost: getEnv("GRPC_HOST", "0.0.0.0"),
		GRPCPort: getIntEnv("GRPC_PORT", 50052),
		HTTPPort: getIntEnv("HTTP_PORT", 8183),
	}
}

func loadEnvFiles() {
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
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
