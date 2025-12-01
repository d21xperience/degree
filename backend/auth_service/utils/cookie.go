package utils

// type CookieCfg struct {
// 	Domain string // contoh ".myapp.com", kosong = host saat ini
// 	Secure bool   // true jika HTTPS
// }

// func SetAuthCookies(w http.ResponseWriter, accessToken, refreshToken string, rememberMe bool) {
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "access_token",
// 		Value:    accessToken,
// 		Path:     "/",
// 		HttpOnly: true,
// 		Secure:   true, // Aktifkan saat pakai HTTPS
// 		MaxAge:   900,  // 15 menit
// 		SameSite: http.SameSiteNoneMode,
// 	})
// 	refreshMaxAge := 60 * 60 * 24 // 1 hari (tanpa Remember Me)
// 	if rememberMe {
// 		refreshMaxAge = 60 * 60 * 24 * 7 // 7 hari (dengan Remember Me)
// 	}
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "refresh_token",
// 		Value:    refreshToken,
// 		Path:     "/",
// 		HttpOnly: true,
// 		Secure:   true,
// 		MaxAge:   refreshMaxAge,
// 		SameSite: http.SameSiteNoneMode,
// 	})
// }

// func ClearAuthCookies(w http.ResponseWriter) {
// 	expired := time.Unix(0, 0) // waktu di masa lalu

// 	for _, name := range []string{"access_token", "refresh_token"} {
// 		http.SetCookie(w, &http.Cookie{
// 			Name:     name,
// 			Value:    "",
// 			Path:     "/",
// 			Expires:  expired,
// 			MaxAge:   -1,
// 			HttpOnly: true,
// 			Secure:   true,
// 			SameSite: http.SameSiteNoneMode,
// 		})
// 	}
// }
