package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func GetPath(folderName, filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get working directory: %v", err)
	}
	contentFile := filepath.Join(wd, folderName, filename)
	return contentFile
}

// Fungsi helper untuk verifikasi file
func VerifyFileAccess(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get working directory: %v", err)
	}
	contentFile := filepath.Join(wd, path)
	if _, err := os.Stat(contentFile); os.IsNotExist(err) {
		return fmt.Errorf("file %s tidak ditemukan", contentFile)
	}

	if _, err := os.ReadFile(contentFile); err != nil {
		return fmt.Errorf("tidak bisa membaca file %s: %v", contentFile, err)
	}

	return nil
}

// Fungsi helper dengan retry mechanism
func ReadFileWithRetry(path string, maxRetries int) ([]byte, error) {
	var err error
	var content []byte

	for i := 0; i < maxRetries; i++ {
		content, err = os.ReadFile(path)
		if err == nil {
			return content, nil
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil, err
}

// bantu reader
func StringReader(s string) *os.File {
	tmp := "/tmp/abi_tmp"
	_ = os.WriteFile(tmp, []byte(s), 0644)
	f, _ := os.Open(tmp)
	return f
}
