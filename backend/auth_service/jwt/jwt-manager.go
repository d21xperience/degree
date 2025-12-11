package jwt

import (
	"auth_service/models"
	"auth_service/utils"
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const accessTokenExpiry = 30 * time.Minute

type Manager struct {
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
	Issuer          string
	JwksURL         string
	Kid             string
}

type Claims struct {
	UserID          string `json:"user_id"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	SekolahTenantId string `json:"sekolah_tenant_id"`
	jwt.RegisteredClaims
}

// func NewManager(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *Manager {
func NewManager() *Manager {
	utils.LoadEnvFiles()
	privateKey, err := loadPrivateKey(utils.GetEnv("JWT_PRIVATE_PATH", "./keys/private.pem"))
	if err != nil {
		panic("tidak dapat membaca private key " + err.Error())
	}
	publicKey, kid, err := ParseKeyFile(utils.GetEnv("JWT_PRIVATE_PATH", "./keys/private.pem"))
	if err != nil {
		panic(err)
	}

	return &Manager{
		PrivateKey:      privateKey,
		PublicKey:       publicKey,
		AccessTokenExp:  accessTokenExpiry,
		RefreshTokenExp: 7 * 24 * time.Hour,
		Issuer:          "sc-app",
		Kid:             kid,
	}
}

func (m *Manager) GenerateTokens(user *models.User) (string, string, error) {
	if m.Kid == "" {
		return "", "", fmt.Errorf("KeyID is not configured")
	}
	// Access Token
	accessClaims := Claims{
		UserID:          fmt.Sprintf("%d", user.ID),
		Email:           user.Email,
		Role:            user.Role,
		SekolahTenantId: fmt.Sprintf("%d", user.SekolahTenantID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.AccessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    m.Issuer,
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	accessToken.Header["kid"] = m.Kid
	accessTokenString, err := accessToken.SignedString(m.PrivateKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshClaims := Claims{
		UserID:          fmt.Sprintf("%d", user.ID),
		Email:           user.Email,
		Role:            user.Role,
		SekolahTenantId: fmt.Sprintf("%d", user.SekolahTenantID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.RefreshTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    m.Issuer,
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	refreshToken.Header["kid"] = m.Kid
	refreshTokenString, err := refreshToken.SignedString(m.PrivateKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (m *Manager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
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

func loadPublicKey(path string) (*rsa.PublicKey, error) {
	// utils.LoadEnvFiles()
	data, _ := os.ReadFile(path)
	key, _ := jwt.ParseRSAPublicKeyFromPEM(data)
	return key, nil
}

// internal/auth/jwt_manager.go
func (j *Manager) VerifyRefreshToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return j.PublicKey, nil // pastikan Anda punya public key untuk refresh token
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// ✅ Verifikasi claim khusus: pastikan ini refresh token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	tokenType, _ := claims["type"].(string)
	if tokenType != "refresh" {
		return nil, errors.New("not a refresh token")
	}

	return token, nil
}
