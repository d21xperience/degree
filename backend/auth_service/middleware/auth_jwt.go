package middleware

import (
	"auth_service/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userClaimsKey contextKey = "userClaims"

// func JWTAuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
// 		path := r.URL.Path

// 		// Hanya path yang mengandung "/auth/" yang memerlukan JWT
// 		if !strings.Contains(path, "/auth/") {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// 🔒 Lakukan validasi JWT hanya untuk endpoint auth
// 		cookie, err := r.Cookie("access_token")
// 		if err != nil || cookie.Value == "" {
// 			http.Error(w, "Unauthorized: no token", http.StatusUnauthorized)
// 			return
// 		}

// 		claims, err := utils.ValidateJWT(cookie.Value)
// 		if err != nil {
// 			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Tambahkan claims ke context agar bisa digunakan di handler
// 		ctx := context.WithValue(r.Context(), userClaimsKey, claims)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		protectedPaths := []string{
			"/api/v1/as/auth/web/me",
			"/api/v1/as/auth/web/refresh",
			// "/api/v1/as/auth/web/logout",
		}

		// Jika path tidak ada dalam daftar yang dilindungi → langsung lewat
		if !isProtectedPath(r.URL.Path, protectedPaths) {
			next.ServeHTTP(w, r)
			return
		}

		// 🔒 Path dilindungi, lakukan validasi token
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

// isProtectedPath memastikan path cocok tepat dengan salah satu prefix yang dilindungi
func isProtectedPath(path string, protected []string) bool {
	// Normalisasi path agar tidak salah match
	cleanPath := strings.TrimSuffix(path, "/")

	for _, p := range protected {
		if cleanPath == p || strings.HasPrefix(cleanPath, p+"/") {
			return true
		}
	}
	return false
}
