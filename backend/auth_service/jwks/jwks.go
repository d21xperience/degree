package jwks

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
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

// ParseKeyFile membaca file PEM (private key PKCS1/PKCS8 / public key PKIX / certificate)
// dan mengembalikan rsa.PublicKey + kid.
func ParseKeyFile(path string) (*rsa.PublicKey, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("read key file: %w", err)
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, "", fmt.Errorf("invalid PEM data")
	}

	// Try different block types
	switch block.Type {
	case "RSA PRIVATE KEY":
		// PKCS#1
		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, "", fmt.Errorf("parse PKCS1 private key: %w", err)
		}
		return pubFromRSAPriv(priv), kidFromRSA(&priv.PublicKey), nil

	case "PRIVATE KEY":
		// PKCS#8 (could be RSA or other)
		k, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, "", fmt.Errorf("parse PKCS8 private key: %w", err)
		}
		switch v := k.(type) {
		case *rsa.PrivateKey:
			return pubFromRSAPriv(v), kidFromRSA(&v.PublicKey), nil
		default:
			return nil, "", fmt.Errorf("unsupported private key type in PKCS8: %T", v)
		}

	case "PUBLIC KEY":
		// SubjectPublicKeyInfo (PKIX)
		pubIface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, "", fmt.Errorf("parse public key: %w", err)
		}
		switch v := pubIface.(type) {
		case *rsa.PublicKey:
			return v, kidFromRSA(v), nil
		default:
			return nil, "", fmt.Errorf("unsupported public key type: %T", v)
		}

	case "CERTIFICATE":
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, "", fmt.Errorf("parse certificate: %w", err)
		}
		switch v := cert.PublicKey.(type) {
		case *rsa.PublicKey:
			return v, kidFromRSA(v), nil
		default:
			return nil, "", fmt.Errorf("unsupported cert public key type: %T", v)
		}

	default:
		// Try generic parse (some keys may be pkcs1/8 but labeled differently)
		// Attempt PKCS1
		if priv, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
			return pubFromRSAPriv(priv), kidFromRSA(&priv.PublicKey), nil
		}
		// Attempt PKCS8
		if k, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
			if v, ok := k.(*rsa.PrivateKey); ok {
				return pubFromRSAPriv(v), kidFromRSA(&v.PublicKey), nil
			}
		}
		// Attempt public key
		if pubIface, err := x509.ParsePKIXPublicKey(block.Bytes); err == nil {
			if v, ok := pubIface.(*rsa.PublicKey); ok {
				return v, kidFromRSA(v), nil
			}
		}
		return nil, "", fmt.Errorf("unsupported PEM block type: %s", block.Type)
	}
}

// helper: ambil public key dari rsa.PrivateKey
func pubFromRSAPriv(pk *rsa.PrivateKey) *rsa.PublicKey {
	if pk == nil {
		return nil
	}
	return &pk.PublicKey
}

// helper: generate kid (SHA-256 truncated) dari modulus
func kidFromRSA(pub *rsa.PublicKey) string {
	sum := sha256.Sum256(pub.N.Bytes())
	// ambil 16 hex chars (8 bytes) untuk kid agar ringkas
	return hex.EncodeToString(sum[:8])
}
