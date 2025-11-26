package utils

import (
	"auth_service/jwks"
	"auth_service/models"
	"context"
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"os"
	"slices"
	"strings"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	LoadEnvFiles()
	data, _ := os.ReadFile(GetEnv("JWT_PUBLIC_PATH", "./keys/public.pem"))
	key, _ := jwt.ParseRSAPublicKeyFromPEM(data)
	return key, nil
}

func GenerateTokenRS256(priv any, user *models.User, sekolahTenant *models.SekolahTenant) (string, int64, error) {
	kid, err := jwks.LoadKID("keys/jwks.json")
	if err != nil {
		return "", 0, status.Errorf(codes.Internal, "cannot load KID: %v", err)
	}
	now := time.Now().UTC()
	exp := now.Add(2 * time.Hour)
	claims := jwt.MapClaims{
		"user_id":           user.ID, //fmt.Sprintf("%d", user.ID),
		"roles":             []string{user.Role},
		"exp":               exp.Unix(),
		"username":          user.Username,
		"asal_sekolah":      sekolahTenant.NamaSekolah,
		"sekolah_tenant_id": sekolahTenant.ID,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t.Header["kid"] = kid
	s, err := t.SignedString(priv)
	return s, exp.Unix(), err
}
