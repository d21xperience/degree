package jwks

import (
	"encoding/json"
	"net/http"
	"time"
)

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func FetchJWKS(url string) (JWKS, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return JWKS{}, err
	}
	defer resp.Body.Close()

	var jwks JWKS
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	return jwks, err
}
