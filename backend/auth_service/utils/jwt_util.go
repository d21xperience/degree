package utils

import (
	"encoding/hex"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key")

//	func GenerateJWT(user any, duration time.Duration) (string, error) {
//		claims := &jwt.StandardClaims{
//			IssuedAt:  time.Now().Unix(),
//			ExpiresAt: time.Now().Add(duration).Unix(),
//		}
//		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//		return token.SignedString(jwtKey)
//	}
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)

		ctx.Next()
	}
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

func GenerateTokenRS256(priv any, userID int64) (string, int64, error) {
	now := time.Now().UTC()
	exp := now.Add(2 * time.Hour)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp.Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	s, err := t.SignedString(priv)
	return s, exp.Unix(), err
}
