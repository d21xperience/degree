package utils

import (
	"context"
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"os"
	"slices"
	"strings"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

type Claims struct {
	UserID          int64    `json:"user_id"`
	Roles           []string `json:"roles"`
	Username        string   `json:"username"`
	AsalSekolah     string   `json:"asal_sekolah"`
	SekolahTenantId int32    `json:"sekolah_tenant_id"`
	jwt.RegisteredClaims
}

func ExtractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata")
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return "", errors.New("no auth header")
	}

	parts := strings.Split(authHeader[0], " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	return parts[1], nil
}

func ParseAndVerify(tokenStr string, pub *rsa.PublicKey) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		// ensure method is RSA
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return pub, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	// Normalize roles: avoid nil slice
	if claims.Roles == nil {
		claims.Roles = []string{}
	}
	return claims, nil
}

// func HasRole(claims *Claims, required string) bool {
// 	for _, r := range claims.Roles {
// 		if r == required {
// 			return true
// 		}
// 	}
// 	return false
// }

func Contains(r []string, target string) bool {
	return slices.Contains(r, target)
}

func GenerateRefreshToken(randRead func([]byte) (int, error)) (string, error) {
	bytes := make([]byte, 32)
	_, err := randRead(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func LoadPrivateKey(path string) (any, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	priv, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func LoadPublicKey() (*rsa.PublicKey, error) {
	data, _ := os.ReadFile("config/public.pem")
	key, _ := jwt.ParseRSAPublicKeyFromPEM(data)
	return key, nil
}
