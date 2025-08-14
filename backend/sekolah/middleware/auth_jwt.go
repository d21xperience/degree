package middleware

// func JWTAuth(secret string) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			var tokenStr string

// 			// 1) Cek header Authorization: Bearer <token>
// 			if h := r.Header.Get("Authorization"); strings.HasPrefix(h, "Bearer ") {
// 				tokenStr = strings.TrimPrefix(h, "Bearer ")
// 			}

// 			// 2) Jika kosong, cek cookie "access_token"
// 			if tokenStr == "" {
// 				if c, err := r.Cookie("access_token"); err == nil {
// 					tokenStr = c.Value
// 				}
// 			}

// 			if tokenStr == "" {
// 				http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 				return
// 			}

// 			_, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
// 				return []byte(secret), nil
// 			})
// 			if err != nil {
// 				http.Error(w, "Invalid token", http.StatusUnauthorized)
// 				return
// 			}

// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// func JWTAuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Allow unauthenticated routes
// 		if r.URL.Path == "/api/v1/as/auth/web/login" || r.URL.Path == "/api/v1/as/auth/web/refresh" || r.URL.Path == "/api/v1/as/auth/web/logout" || r.URL.Path == "/api/v1/as/sekolah" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// Validasi token
// 		cookie, err := r.Cookie("access_token")
// 		if err != nil || cookie.Value == "" {
// 			http.Error(w, "Unauthorized: no token", http.StatusUnauthorized)
// 			return
// 		}

// 		_, err = utils.ValidateJWT(cookie.Value)
// 		if err != nil {
// 			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
