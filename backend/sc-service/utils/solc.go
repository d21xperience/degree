package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CompileSolidityFile(solPath string) (abiPath string, binPath string, err error) {
	outputDir := filepath.Dir(solPath)

	cmd := exec.Command("solc", "--abi", "--bin", solPath, "-o", outputDir)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("compile error: %w", err)
	}

	baseName := strings.TrimSuffix(filepath.Base(solPath), ".sol")
	abiPath = filepath.Join(outputDir, baseName+".abi")
	binPath = filepath.Join(outputDir, baseName+".bin")
	return abiPath, binPath, nil
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
