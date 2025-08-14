package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CompileSolidityFile(srcPath string) (string, string, error) {
	dir := filepath.Dir(srcPath)
	base := filepath.Base(srcPath)
	solName := strings.TrimSuffix(base, filepath.Ext(base)) // misal: ijazah_v3

	// Jalankan solc
	cmd := exec.Command("solc", "--abi", "--bin", srcPath, "-o", dir, "--overwrite")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", fmt.Errorf("kompilasi gagal: %v\nOutput: %s", err, string(output))
	}

	// Temukan file .abi dan .bin yang baru saja dihasilkan
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", "", fmt.Errorf("gagal membaca direktori: %w", err)
	}

	var abiFile, binFile string
	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".abi") && abiFile == "" {
			abiFile = name
		}
		if strings.HasSuffix(name, ".bin") && binFile == "" {
			binFile = name
		}
	}

	if abiFile == "" || binFile == "" {
		return "", "", errors.New("file ABI atau BIN tidak ditemukan setelah kompilasi")
	}

	// Rename file menjadi nama file .sol
	newAbiPath := filepath.Join(dir, solName+".abi")
	newBinPath := filepath.Join(dir, solName+".bin")

	err = os.Rename(filepath.Join(dir, abiFile), newAbiPath)
	if err != nil {
		return "", "", fmt.Errorf("gagal rename ABI: %w", err)
	}

	err = os.Rename(filepath.Join(dir, binFile), newBinPath)
	if err != nil {
		return "", "", fmt.Errorf("gagal rename BIN: %w", err)
	}

	return newAbiPath, newBinPath, nil
}

func GetSolcVersion() (string, error) {
	cmd := exec.Command("solc", "--version")

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Version:") {
			return strings.TrimSpace(line), nil
		}
	}
	return "Unknown version", nil
}

// func fileExistsAndNotEmpty(path string) bool {
// 	info, err := os.Stat(path)
// 	if os.IsNotExist(err) || info.Size() == 0 {
// 		return false
// 	}
// 	return true
// }
