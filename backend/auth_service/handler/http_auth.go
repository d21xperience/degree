package handler

import (
	"auth_service/config"
	"auth_service/models"
	"auth_service/repositories"
	"auth_service/services"
	"auth_service/utils"
	"encoding/json"
	"net/http"
	"time"
)

type SekolahTenant struct {
	NamaSekolah string `json:"namaSekolah"`
	Npsn        string `json:"Npsn"`
	EnkripID    string `json:"enkripId"`
}

type UserLoggedIn struct {
	ID              uint64 `json:"id"`
	SekolahTenantID uint32 `json:"sekolahTenantId"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Role            string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status        bool           `json:"status"`
	Message       string         `json:"message,omitempty"`
	User          UserLoggedIn   `json:"user"`
	SekolahTenant *SekolahTenant `json:"sekolahTenant"`
}

type handlerHTTP struct {
	Auth        services.AuthService
	repoSekolah repositories.SekolahRepository
}

func NewHandlerHttp() *handlerHTTP {
	repoAuth := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthService(repoAuth)
	repoSekolah := repositories.NewSekolahRepository(config.DB)
	return &handlerHTTP{
		Auth:        authService,
		repoSekolah: repoSekolah,
	}
}
 
func (h handlerHTTP) HandlerLoginHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.Password == "" || (req.Email == "" && req.Username == "") {
			http.Error(w, "Email/username dan password wajib diisi", http.StatusBadRequest)
			return
		}

		user, err := func() (*models.User, error) {
			if req.Email != "" {
				return h.Auth.Login(req.Email, req.Password)
			}
			return h.Auth.Login(req.Username, req.Password)
		}()
		if err != nil {
			http.Error(w, "Email/username atau password salah", http.StatusUnauthorized)
			return
		}

		userLoggedin := UserLoggedIn{
			ID:              user.ID,
			Username:        user.Username,
			Email:           user.Email,
			Role:            user.Role,
			SekolahTenantID: user.SekolahTenantID,
		}
		var sekolahModel *models.SekolahTenant
		var SekolahTenant SekolahTenant
		if user.Role != "superadmin" {
			sekolahModel, err = h.repoSekolah.GetSekolahByTenantId(user.SekolahTenantID)
			if err != nil {
				http.Error(w, "Gagal ambil data sekolah", http.StatusUnauthorized)
				return
			}

			SekolahTenant.NamaSekolah = sekolahModel.NamaSekolah
			SekolahTenant.Npsn = sekolahModel.NPSN
			SekolahTenant.EnkripID = sekolahModel.EnkripID
		}

		access, _ := utils.GenerateJWT(user, 15*time.Minute)
		refresh, _ := utils.GenerateJWT(user, 7*24*time.Hour)
		utils.SetAuthCookies(w, access, refresh)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(LoginResponse{
			Status:        true,
			Message:       "Login berhasil",
			User:          userLoggedin,
			SekolahTenant: &SekolahTenant,
		})
	}
}

func (h handlerHTTP) HandlerRefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("refresh_token")
		if err != nil || c.Value == "" {
			http.Error(w, "Refresh token missing", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(c.Value)
		if err != nil {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			return
		}

		// (Opsional) cek token rotation di DB / Redis di sini

		user, err := h.Auth.GetUserByID(claims.UserID)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Generate token baru
		access, _ := utils.GenerateJWT(user, 15*time.Minute)
		refresh, _ := utils.GenerateJWT(user, 7*24*time.Hour)

		utils.SetAuthCookies(w, access, refresh)

		json.NewEncoder(w).Encode(map[string]any{
			"status":  true,
			"message": "refreshed",
		})
	}
}

func (h handlerHTTP) HandlerLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ClearAuthCookies(w)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  true,
			"message": "logged out",
		})
	}
}
