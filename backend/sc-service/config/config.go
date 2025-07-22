package config

import (
	"sc-service/utils"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func LoadConfig() Config {
	_ = godotenv.Load(".env") // Cari file .env spesifik
	_ = godotenv.Load()       // Fallback ke default .env
	return Config{
		User:            utils.GetEnvWithSecretFallback("SCDB_USER", "postgres"),
		Password:        utils.GetEnvWithSecretFallback("SCDB_PASSWORD", "postgres"),
		Host:            utils.GetEnv("SCDB_HOST", "localhost"),
		Port:            utils.GetEnvAsInt("SCDB_PORT", 5444),
		DBName:          utils.GetEnv("SCDB_DB", "dbsc"),
		MaxIdleConns:    utils.GetEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		MaxOpenConns:    utils.GetEnvAsInt("DB_MAX_OPEN_CONNS", 100),
		ConnMaxLifetime: time.Duration(utils.GetEnvAsInt("DB_CONN_MAX_LIFETIME_MIN", 30)) * time.Minute,
	}
}
