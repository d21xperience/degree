package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sc-service/utils"
)

// func HandlerCompileContractHTTP() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Validasi method
// 		if r.Method != http.MethodPost {
// 			http.Error(w, "Hanya POST diizinkan", http.StatusMethodNotAllowed)
// 			return
// 		}

// 		// Parse multipart form
// 		err := r.ParseMultipartForm(10 << 20) // Max 10MB
// 		if err != nil {
// 			http.Error(w, "Gagal membaca form", http.StatusBadRequest)
// 			return
// 		}

// 		file, handler, err := r.FormFile("file")
// 		if err != nil {
// 			http.Error(w, "File .sol tidak ditemukan", http.StatusBadRequest)
// 			return
// 		}
// 		defer file.Close()

// 		// Simpan file sementara
// 		tempDir := "./contracts"
// 		_ = os.MkdirAll(tempDir, os.ModePerm)
// 		dstPath := filepath.Join(tempDir, handler.Filename)
// 		dst, _ := os.Create(dstPath)
// 		defer dst.Close()
// 		_, _ = io.Copy(dst, file)

// 		// Compile
// 		abiPath, binPath, err := utils.CompileSolidityFile(dstPath)
// 		if err != nil {
// 			http.Error(w, "Compile gagal: "+err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		log.Printf("%v %v", abiPath, binPath)
// 		abiContent, _ := os.ReadFile(abiPath)
// 		binContent, _ := os.ReadFile(binPath)

//			// Response
//			w.Header().Set("Content-Type", "application/json")
//			json.NewEncoder(w).Encode(map[string]any{
//				"status":   true,
//				"message":  "Compile berhasil",
//				"abi":      string(abiContent),
//				"bytecode": string(binContent),
//			})
//		}
//	}
func HandlerCompileContractHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validasi method
		if r.Method != http.MethodPost {
			http.Error(w, "Hanya POST diizinkan", http.StatusMethodNotAllowed)
			return
		}

		// Parse multipart form
		err := r.ParseMultipartForm(10 << 20) // Max 10MB
		if err != nil {
			http.Error(w, "Gagal membaca form: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Ambil file dari form
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "File .sol tidak ditemukan: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Buat direktori kontrak jika belum ada
		tempDir := "./contracts"
		if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
			http.Error(w, "Gagal membuat direktori: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Simpan file sementara
		dstPath := filepath.Join(tempDir, handler.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Gagal membuat file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Gagal menyalin file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Compile kontrak
		abiPath, binPath, err := utils.CompileSolidityFile(dstPath)
		if err != nil {
			http.Error(w, "Compile gagal: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Baca konten hasil compile
		abiContent, err := os.ReadFile(abiPath)
		if err != nil {
			http.Error(w, "Gagal membaca file ABI: "+err.Error(), http.StatusInternalServerError)
			return
		}

		binContent, err := os.ReadFile(binPath)
		if err != nil {
			http.Error(w, "Gagal membaca file bytecode: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Bersihkan file temporary
		defer func() {
			os.Remove(dstPath)
			os.Remove(abiPath)
			os.Remove(binPath)
		}()

		// Response
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{
			"status":   "true",
			"message":  "Compile berhasil",
			"abi":      string(abiContent),
			"bytecode": string(binContent),
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Gagal mengencode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
