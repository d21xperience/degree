package handlers

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

// 		abiContent, _ := os.ReadFile(abiPath)
// 		binContent, _ := os.ReadFile(binPath)

// 		// Response
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]any{
// 			"message":  "Compile berhasil",
// 			"abi":      string(abiContent),
// 			"bytecode": string(binContent),
// 		})
// 	}
// }
