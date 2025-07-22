package middleware

import (
	"auth_service/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userClaimsKey contextKey = "userClaims"

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedPaths := []string{
			"/api/v1/as/auth/web/me",
			"/api/v1/as/auth/web/refresh",
			"/api/v1/as/auth/web/login",
			"/api/v1/as/auth/web/logout",
			"/api/v1/as/sekolah",
		}

		if containsPathPrefix(allowedPaths, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("access_token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized: no token", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func containsPathPrefix(paths []string, target string) bool {
	for _, path := range paths {
		if strings.HasPrefix(target, path) {
			return true
		}
	}
	return false
}
