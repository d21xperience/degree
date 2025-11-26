package jwks

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"math/big"
)

func JWKToRSA(jwk JWK) (*rsa.PublicKey, error) {
	nb, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, err
	}
	eb, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, err
	}

	e := int(new(big.Int).SetBytes(eb).Int64())
	pub := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}
	return pub, nil
}

func FindKey(jwks JWKS, kid string) (JWK, error) {
	for _, k := range jwks.Keys {
		if k.Kid == kid {
			return k, nil
		}
	}
	return JWK{}, errors.New("kid not found")
}
