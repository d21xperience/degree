package main

import (
	"sekolah/config"
	"sekolah/server"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()
	// Inisialisasi database
	config.InitDatabase(cfg)
	// Start redis
	config.InitRedisClient(cfg)
	// Migrasi model
	// config.DB.AutoMigrate(&models.SekolahTabelTenant{}, &models.SchemaLog{})
	server.StartServer()
}
