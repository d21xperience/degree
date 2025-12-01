package jwks

import (
	"auth_service/jwt"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
)

func JWKSHandler(jwtManager *jwt.Manager) func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

		n := base64.RawURLEncoding.EncodeToString(jwtManager.PublicKey.N.Bytes())
		e := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(jwtManager.PublicKey.E)).Bytes())

		jwk := JWK{
			Kid: jwtManager.Kid,
			Kty: "RSA",
			Alg: "RS256",
			Use: "sig",
			N:   n,
			E:   e,
		}

		jwks := JWKS{Keys: []JWK{jwk}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jwks)
	}
}
