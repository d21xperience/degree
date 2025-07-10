package handlers

import (
	"encoding/json"
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"
)

func HandlerIPFSHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // max 10MB
		if err != nil {
			http.Error(w, `{"error":"Invalid form data"}`, http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, `{"error":"Invalid file"}`, http.StatusBadRequest)
			return
		}
		defer file.Close()

		sh := shell.NewShell("localhost:5001")
		cid, err := sh.Add(file)
		if err != nil {
			http.Error(w, `{"error":"Failed to upload to IPFS"}`, http.StatusInternalServerError)
			return
		}

		// newMetadata := models.IPFSMetadata{
		// 	IPFSCID:    cid,
		// 	Size:       int(handler.Size),
		// 	FileName:   handler.Filename,
		// 	FileType:   handler.Header.Get("Content-Type"), // atau baca sebagian konten
		// 	UploadedAt: time.Now(),
		// }
		// newMetadata := models.IPFSMetadata{
		// 	IPFSCID:    cid,
		// 	Size:       int(handler.Size),
		// 	FileName:   handler.Filename,
		// 	FileType:   handler.Header.Get("Content-Type"), // atau baca sebagian konten
		// 	UploadedAt: time.Now(),
		// }

		// if err := ac.DB.Create(&newMetadata).Error; err != nil {
		// 	http.Error(w, `{"error":"Failed to save metadata"}`, http.StatusInternalServerError)
		// 	return
		// }

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"cid": cid})

	}
}
