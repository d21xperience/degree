package jwt

import (
	"auth_service/models"
	"auth_service/utils"
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
	Issuer          string
	JwksURL         string
}

type Claims struct {
	// UserID          int64  `json:"user_id"`
	Role            string `json:"role"`
	Username        string `json:"username"`
	AsalSekolah     string `json:"asal_sekolah"`
	SekolahTenantId int32  `json:"sekolah_tenant_id"`
	jwt.RegisteredClaims
}

// func NewManager(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *Manager {
func NewManager() *Manager {
	utils.LoadEnvFiles()
	privateKey, err := loadPrivateKey(utils.GetEnv("JWT_PRIVATE_PATH", "./keys/private.pem"))
	if err != nil {
		panic("tidak dapat membaca private key " + err.Error())
	}

	publicKey, err := loadPublicKey(utils.GetEnv("JWT_PUBLIC_PATH", "./keys/public.pem"))
	if err != nil {
		panic("tidak dapat membaca public key " + err.Error())
	}

	return &Manager{
		PrivateKey:      privateKey,
		PublicKey:       publicKey,
		AccessTokenExp:  15 * time.Minute,
		RefreshTokenExp: 7 * 24 * time.Hour,
		Issuer:          "sc-app",
	}
}

func (m *Manager) GenerateTokens(user *models.User, sekolahTenant *models.SekolahTenant) (string, string, error) {
	// Access Token
	accessClaims := Claims{
		// UserID:          user.ID,
		Role:            user.Role,
		Username:        user.Username,
		AsalSekolah:     sekolahTenant.NamaSekolah,
		SekolahTenantId: sekolahTenant.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.AccessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    m.Issuer,
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(m.PrivateKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshClaims := Claims{
		Role:            user.Role,
		Username:        user.Username,
		AsalSekolah:     sekolahTenant.NamaSekolah,
		SekolahTenantId: sekolahTenant.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.RefreshTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    m.Issuer,
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
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
