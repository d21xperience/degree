package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CompileSolidityFile(solPath string) (abiPath string, binPath string, err error) {
	baseName := strings.TrimSuffix(filepath.Base(solPath), ".sol")
	outputDir := filepath.Dir(solPath)
	fmt.Println(outputDir)

	abiPath = filepath.Join(outputDir, baseName+".abi")
	binPath = filepath.Join(outputDir, baseName+".bin")
	// Path file dalam container

	// Path file dalam container
	containerSolPath := "/sources/" + filepath.Base(solPath)

	// Jalankan docker run untuk compile
	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/sources", outputDir),
		"ethereum/solc:0.8.20", // atau :stable jika ingin latest stable
		containerSolPath,
		"--abi", "--bin",
		"--output-dir", "/sources",
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", "", fmt.Errorf("compile error: %s", stderr.String())
	}

	// Validasi file berhasil dibuat
	if !fileExistsAndNotEmpty(abiPath) || !fileExistsAndNotEmpty(binPath) {
		return "", "", fmt.Errorf("ABI atau bytecode tidak ditemukan")
	}

	return abiPath, binPath, nil
}

func GetSolcVersion() (string, error) {
	cmd := exec.Command(
		"docker", "run", "--rm",
		"ethereum/solc:0.8.30-alpine",
		"solc", "--version",
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	log.Println("🛠 Menjalankan perintah:", cmd.String()) // ⬅️ Tambah log
	err := cmd.Run()
	log.Println("🔎 Output:", out.String()) // ⬅️ Tambah log
	if err != nil {
		return "", fmt.Errorf("solc version error: %s", out.String())
	}
	// if err := cmd.Run(); err != nil {
	// 	return "", fmt.Errorf("solc version error: %s", out.String())
	// }

	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Version:") {
			return strings.TrimSpace(line), nil
		}
	}
	return "Versi tidak ditemukan", nil
}

func fileExistsAndNotEmpty(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info.Size() == 0 {
		return false
	}
	return true
}
