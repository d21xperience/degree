/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
// package main

// func main() {
// 	// cmd.Execute()

// }
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Fungsi utama untuk mengirim request dengan Bearer token
func main() {
	// Load .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("No .env file found, using default env.")
	// }
	// cmd.Execute()
	tes := buatSingkatan("Akuntansi")
	fmt.Println(tes)
}

func buatSingkatan(nama string) string {
	var kataYangDiabaikan = map[string]bool{
		"dan":    true,
		"dari":   true,
		"ke":     true,
		"yang":   true,
		"untuk":  true,
		"dengan": true,
		"di":     true,
	}

	// Singkatan khusus (override)
	var singkatanKhusus = map[string]string{
		"akuntansi": "AK",
	}

	// Cek override khusus
	normalisasi := strings.ToLower(strings.TrimSpace(nama))
	if val, ok := singkatanKhusus[normalisasi]; ok {
		return val
	}

	kata := strings.Fields(normalisasi)
	singkatan := ""

	for _, k := range kata {
		if kataYangDiabaikan[k] {
			continue
		}
		for _, r := range k {
			if unicode.IsLetter(r) {
				singkatan += strings.ToUpper(string(r))
				break
			}
		}
	}

	return singkatan
}
