package jwks

import (
	"encoding/json"
	"fmt"
	"os"
)

// Struct JWK/JWKS
type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

func LoadKID(jwksPath string) (string, error) {
	b, err := os.ReadFile(jwksPath)
	if err != nil {
		return "", err
	}

	var file JWKS
	if err := json.Unmarshal(b, &file); err != nil {
		return "", err
	}

	if len(file.Keys) == 0 {
		return "", fmt.Errorf("no keys found in JWKS file")
	}

	return file.Keys[0].Kid, nil
}
