package http_handler

import (
	"net/http"
	"time"
)

type CookieConfig struct {
	Domain   string
	Secure   bool
	SameSite http.SameSite
}

type CookieHandler struct {
	config CookieConfig
}

func NewCookieHandler(config CookieConfig) *CookieHandler {
	return &CookieHandler{config: config}
}

func (h *CookieHandler) SetTokens(w http.ResponseWriter, accessToken, refreshToken string) {
	// Set access token cookie (short-lived)
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		Domain:   h.config.Domain,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   h.config.Secure,
		SameSite: h.config.SameSite,
	})

	// Set refresh token cookie (long-lived)
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Domain:   h.config.Domain,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   h.config.Secure,
		SameSite: h.config.SameSite,
	})
}

func (h *CookieHandler) ClearTokens(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Domain:   h.config.Domain,
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Secure:   h.config.Secure,
		SameSite: h.config.SameSite,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Domain:   h.config.Domain,
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Secure:   h.config.Secure,
		SameSite: h.config.SameSite,
	})
}

func (h *CookieHandler) GetTokenFromCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
