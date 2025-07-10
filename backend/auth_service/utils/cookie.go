package utils

import "net/http"

type CookieCfg struct {
	Domain string // contoh ".myapp.com", kosong = host saat ini
	Secure bool   // true jika HTTPS
}

// SetAuthCookies menulis access & refresh token.
// func SetAuthCookies(w http.ResponseWriter, access, refresh string, cfg CookieCfg) {
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "access_token",
// 		Value:    access,
// 		Path:     "/",
// 		Domain:   cfg.Domain,
// 		HttpOnly: true,
// 		Secure:   cfg.Secure,
// 		MaxAge:   15 * 60,
// 		SameSite: http.SameSiteStrictMode,
// 	})
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "refresh_token",
// 		Value:    refresh,
// 		Path:     "/",
// 		Domain:   cfg.Domain,
// 		HttpOnly: true,
// 		Secure:   cfg.Secure,
// 		MaxAge:   7 * 24 * 60 * 60,
// 		SameSite: http.SameSiteStrictMode,
// 	})
// }

func SetAuthCookies(w http.ResponseWriter, accessToken, refreshToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // Aktifkan saat pakai HTTPS
		MaxAge:   15 * 60,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   7 * 24 * 60 * 60,
		SameSite: http.SameSiteStrictMode,
	})
}

// ClearAuthCookies menghapus keduanya.
// func ClearAuthCookies(w http.ResponseWriter, cfg CookieCfg) {
// 	for _, name := range []string{"access_token", "refresh_token"} {
// 		http.SetCookie(w, &http.Cookie{
// 			Name:     name,
// 			Value:    "",
// 			Path:     "/",
// 			Domain:   cfg.Domain,
// 			MaxAge:   -1,
// 			HttpOnly: true,
// 			Secure:   cfg.Secure,
// 			SameSite: http.SameSiteStrictMode,
// 		})
// 	}
// }
func ClearAuthCookies(w http.ResponseWriter) {
	for _, name := range []string{"access_token", "refresh_token"} {
		http.SetCookie(w, &http.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
	}
}
