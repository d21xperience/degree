// server/handler_http.go

package http_handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	pb "auth_service/generated" // sesuaikan path

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HTTPHandler struct {
	AuthClient pb.AuthServiceClient // gRPC client
}

// HandlerLoginHTTP — override /login untuk set HTTP-only cookie
func (h *HTTPHandler) HandlerLoginHTTP() func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		var req pb.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// Panggil gRPC service
		resp, err := h.AuthClient.Login(r.Context(), &req)
		if err != nil {
			// ✅ Parse gRPC error dengan status
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.Unauthenticated:
					http.Error(w, "invalid credentials", http.StatusUnauthorized) // 401
				case codes.InvalidArgument:
					http.Error(w, "invalid request", http.StatusBadRequest) // 400
				case codes.PermissionDenied:
					http.Error(w, "access denied", http.StatusForbidden) // 403
				case codes.Internal:
					// Log detail internal error (jangan tampilkan ke client)
					log.Printf("Internal gRPC error: %v", s.Message())
					http.Error(w, "internal server error", http.StatusInternalServerError) // 500
				case codes.Unavailable:
					log.Printf("gRPC service unavailable: %v", s.Message())
					http.Error(w, "service temporarily unavailable", http.StatusServiceUnavailable) // 503
				default:
					log.Printf("Unhandled gRPC error: code=%v, msg=%q", s.Code(), s.Message())
					http.Error(w, "request failed", http.StatusInternalServerError)
				}
			} else {
				// Bukan error gRPC (misal context canceled, network timeout, dll)
				log.Printf("Non-gRPC error: %v", err)
				http.Error(w, "request failed", http.StatusInternalServerError)
			}
			return
		}

		// ✅ Set cookies (HTTP-only, Secure, SameSite)
		setTokenCookie(w, "access_token", resp.AccessToken, 15*time.Minute)
		setTokenCookie(w, "refresh_token", resp.RefreshToken, 7*24*time.Hour)

		// ✅ Hanya kirim data user, bukan token (aman)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  true,
			"message": "logged on",
			"user_id": resp.UserId,
			"email":   resp.Email,
			// "email":   resp.Email,
		})
	}
}

// HandlerRefreshHTTP — override /refresh untuk rotate & set ulang cookie
func (h *HTTPHandler) HandlerRefreshHTTP() func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		// Perhatikan: refresh_token dikirim via cookie, jadi body bisa kosong
		// Tapi gRPC method tetap butuh request → kirim dummy
		resp, err := h.AuthClient.RefreshToken(r.Context(), &pb.RefreshRequest{})
		if err != nil {
			http.Error(w, "invalid or expired refresh token", http.StatusUnauthorized)
			return
		}

		// ✅ ROTASI: set access & refresh token BARU
		setTokenCookie(w, "access_token", resp.AccessToken, 15*time.Minute)
		setTokenCookie(w, "refresh_token", resp.RefreshToken, 7*24*time.Hour)

		// ✅ Kirim user info (opsional)
		json.NewEncoder(w).Encode(map[string]string{
			"user_id": resp.AccessToken,
			// "email":   resp.Email,
		})
	}
}

// HandlerLogoutHTTP — hapus cookies
func (h *HTTPHandler) HandlerLogoutHTTP() func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		clearTokenCookie(w, "access_token")
		clearTokenCookie(w, "refresh_token")
		w.WriteHeader(http.StatusNoContent)
	}
}

// ---- Helper Cookie ----

func setTokenCookie(w http.ResponseWriter, name, token string, maxAge time.Duration) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		Path:     "/",
		MaxAge:   int(maxAge.Seconds()),
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production", // true di HTTPS
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}

func clearTokenCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // hapus
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		SameSite: http.SameSiteStrictMode,
	})
}
