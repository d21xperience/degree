package main

import (
	"sc-service/config"
	"sc-service/server"
)

func main() {
	// Debugging
	// pwd, _ := os.Getwd()
	// log.Printf("Current working dir: %s", pwd)
	// data, _ := os.ReadFile(".env")
	// log.Printf("Contents of .env:\n%s", string(data))

	// Load konfigurasi database
	cfg := config.LoadConfig()
	// Inisialisasi database
	config.InitDatabase(cfg)
	server.StartServer()
}
