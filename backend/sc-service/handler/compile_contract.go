package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sc-service/utils"
)

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
			utils.RespondWithError(w, http.StatusInternalServerError, "ABI tidak valid: "+err.Error())
			return
		}

		// Pastikan file ada dan bisa diakses
		if err := utils.VerifyFileAccess(abiPath); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "ABI tidak valid: "+err.Error())
			return
		}

		if err := utils.VerifyFileAccess(binPath); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Bytecode tidak valid: "+err.Error())
			return
		}

		// Baca konten dengan retry jika diperlukan
		abiContent, err := utils.ReadFileWithRetry(abiPath, 3)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Gagal membaca ABI setelah 3 percobaan: "+err.Error())
			return
		}

		binContent, err := utils.ReadFileWithRetry(binPath, 3)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Gagal membaca bytecode setelah 3 percobaan: "+err.Error())
			return
		}
		log.Println("=== Bytecode Content ===")
		log.Println(string(binContent))
		// // Bersihkan file temporary
		// defer func() {
		// 	os.Remove(dstPath)
		// 	os.Remove(abiPath)
		// 	os.Remove(binPath)
		// }()

		// Response
		w.Header().Set("Content-Type", "application/json")
		response := map[string]any{
			"status":            true,
			"message":           "Compile berhasil",
			"abi_filename":      filepath.Base(abiPath),
			"bytecode_filename": filepath.Base(binPath),
			"abi":               string(abiContent),
			"bytecode":          string(binContent),
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Gagal mengencode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
