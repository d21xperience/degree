package jwks

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
)

//	func JWKSHandler(pub *rsa.PublicKey, kid string) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			j, err := BuildJWKS(pub, kid)
//			if err != nil {
//				http.Error(w, "cannot build JWKS", 500)
//				return
//			}
//			w.Header().Set("Content-Type", "application/json")
//			w.Write(j)
//		}
//	}
func JWKSHandler(pub *rsa.PublicKey, kid string) func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

		n := base64.RawURLEncoding.EncodeToString(pub.N.Bytes())
		e := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(pub.E)).Bytes())

		jwk := JWK{
			Kid: kid,
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
