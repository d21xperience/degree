package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	secretBytes := make([]byte, 32) // 256-bit
	_, err := rand.Read(secretBytes)
	if err != nil {
		fmt.Println("Gagal generate secret:", err)
		os.Exit(1)
	}

	secret := hex.EncodeToString(secretBytes)
	fmt.Println("Generated JWT_SECRET:", secret)

	// Tulis ke file .env
	f, err := os.Create(".env")
	if err != nil {
		fmt.Println("Gagal menulis file .env:", err)
		os.Exit(1)
	}
	defer f.Close()

	f.WriteString("JWT_SECRET=" + secret + "\n")
	fmt.Println("JWT_SECRET ditulis ke .env")
}
