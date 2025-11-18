package config

import (
	"auth_service/utils"
	"time"
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
	utils.LoadEnvFiles()
	return Config{
		// Redis
		RedisHost:     utils.GetEnv("REDIS_HOST", "localhost"),
		RedisPort:     utils.GetIntEnv("REDIS_PORT", 6379),
		RedisPassword: utils.GetEnv("REDIS_PASSWORD", ""),
		RedisDBName:   utils.GetIntEnv("REDIS_DB", 0),
		// Database
		DBHost:     utils.GetEnv("DB_HOST", "localhost"),
		DBPort:     utils.GetIntEnv("DB_PORT", 5432),
		DBUser:     utils.GetEnv("DB_USER", "postgres"),
		DBPassword: utils.GetEnv("DB_PASSWORD", ""),
		DBName:     utils.GetEnv("DB_NAME", "authdb"),
		// Connection Pool
		MaxIdleConns:    utils.GetIntEnv("MAX_IDLE_CONNS", 10),
		MaxOpenConns:    utils.GetIntEnv("MAX_OPEN_CONNS", 100),
		ConnMaxLifetime: time.Duration(utils.GetIntEnv("CONN_MAX_LIFETIME_MINUTES", 30)) * time.Minute,
		// App
		GRPCHost: utils.GetEnv("GRPC_HOST", "localhost"),
		GRPCPort: utils.GetIntEnv("GRPC_PORT", 50051),
		HTTPPort: utils.GetIntEnv("HTTP_PORT", 8182),
	}
}
