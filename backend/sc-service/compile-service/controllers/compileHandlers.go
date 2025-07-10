package controllers

import (
	"compile-service/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func CompileHandler(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file missing"})
		return
	}
	if !strings.HasSuffix(header.Filename, ".sol") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file harus berekstensi .sol"})
		return
	}

	file, err := header.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open uploaded file"})
		return
	}
	defer file.Close()

	// Simpan file sementara
	tempDir := "./sources"
	_ = os.MkdirAll(tempDir, os.ModePerm)
	dstPath := filepath.Join(tempDir, header.Filename)
	dst, _ := os.Create(dstPath)
	defer dst.Close()
	_, _ = io.Copy(dst, file)

	// Compile
	abiPath, binPath, err := utils.CompileSolidityFile(dstPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "compile error", "detail": err.Error()})
		return
	}

	abiContent, _ := os.ReadFile(abiPath)
	binContent, _ := os.ReadFile(binPath)
	c.JSON(http.StatusOK, gin.H{
		"abi":      string(abiContent),
		"bytecode": string(binContent),
		"message":  "Compile berhasil",
	})
}

func VersionCompilerHandler(c *gin.Context) {
	cek, err := utils.GetSolcVersion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "tidak dapat mengambil versi"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil menggambil versi",
		"versi":   cek,
	})
}
