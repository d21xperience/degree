package handler

// type SekolahTenant struct {
// 	NamaSekolah string `json:"namaSekolah"`
// 	Npsn        string `json:"Npsn"`
// 	EnkripID    string `json:"enkripId"`
// }

// type UserLoggedIn struct {
// 	ID              int64  `json:"id"`
// 	SekolahTenantID uint32 `json:"sekolahTenantId"`
// 	Username        string `json:"username"`
// 	Email           string `json:"email"`
// 	Role            string `json:"role"`
// }

// type LoginRequest struct {
// 	Email    string `json:"email"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type LoginResponse struct {
// 	Status        bool           `json:"status"`
// 	Message       string         `json:"message,omitempty"`
// 	User          UserLoggedIn   `json:"user"`
// 	SekolahTenant *SekolahTenant `json:"sekolahTenant"`
// }

// type handlerHTTP struct {
// 	Auth        services.AuthService
// 	repoSekolah repositories.SekolahRepository
// }

// func NewHandlerHttp() *handlerHTTP {
// 	repoAuth := repositories.NewUserRepository(config.DB)
// 	authService := services.NewAuthService(repoAuth)
// 	repoSekolah := repositories.NewSekolahRepository(config.DB)
// 	return &handlerHTTP{
// 		Auth:        authService,
// 		repoSekolah: repoSekolah,
// 	}
// }

// func (h handlerHTTP) HandlerLoginHTTP() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req LoginRequest
// 		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 			http.Error(w, "Invalid request body", http.StatusBadRequest)
// 			return
// 		}
// 		if req.Password == "" || (req.Email == "" && req.Username == "") {
// 			http.Error(w, "Email/username dan password wajib diisi", http.StatusBadRequest)
// 			return
// 		}

// 		user, err := func() (*models.User, error) {
// 			if req.Email != "" {
// 				return h.Auth.Login(req.Email, req.Password)
// 			}
// 			return h.Auth.Login(req.Username, req.Password)
// 		}()
// 		if err != nil {
// 			http.Error(w, "Email/username atau password salah", http.StatusUnauthorized)
// 			return
// 		}

// 		userLoggedin := UserLoggedIn{
// 			ID:              user.ID,
// 			Username:        user.Username,
// 			Email:           user.Email,
// 			Role:            user.Role,
// 			SekolahTenantID: user.SekolahTenantID,
// 		}
// 		var sekolahModel *models.SekolahTenant
// 		var SekolahTenant SekolahTenant
// 		if user.Role != "superadmin" {
// 			sekolahModel, err = h.repoSekolah.GetSekolahByTenantId(user.SekolahTenantID)
// 			if err != nil {
// 				http.Error(w, "Gagal ambil data sekolah", http.StatusUnauthorized)
// 				return
// 			}

// 			SekolahTenant.NamaSekolah = sekolahModel.NamaSekolah
// 			SekolahTenant.Npsn = sekolahModel.NPSN
// 			SekolahTenant.EnkripID = sekolahModel.EnkripID
// 		}
// 		rememberMe := r.FormValue("remember_me") == "true"
// 		// 2. Generate tokens
// 		access, _ := utils.GenerateJWT(user, 15*time.Minute)
// 		refresh, _ := utils.GenerateJWT(user, 7*24*time.Hour)

// 		// 3. Set cookies
// 		utils.SetAuthCookies(w, access, refresh, rememberMe)

// 		w.Header().Set("Content-Type", "application/json")
// 		_ = json.NewEncoder(w).Encode(LoginResponse{
// 			Status:        true,
// 			Message:       "Login berhasil",
// 			User:          userLoggedin,
// 			SekolahTenant: &SekolahTenant,
// 		})
// 	}
// }

// func (h handlerHTTP) HandlerRefreshToken() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		c, err := r.Cookie("refresh_token")
// 		if err != nil || c.Value == "" {
// 			http.Error(w, "Refresh token missing", http.StatusUnauthorized)
// 			return
// 		}

// 		claims, err := utils.ValidateJWT(c.Value)
// 		if err != nil {
// 			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
// 			return
// 		}

// 		// (Opsional) cek token rotation di DB / Redis di sini

// 		user, err := h.Auth.GetUserByID(claims.UserID)
// 		if err != nil {
// 			http.Error(w, "User not found", http.StatusUnauthorized)
// 			return
// 		}

// 		// Generate token baru
// 		access, _ := utils.GenerateJWT(user, 15*time.Minute)
// 		refresh, _ := utils.GenerateJWT(user, 7*24*time.Hour)

// 		rememberMe := r.FormValue("remember_me") == "true"
// 		utils.SetAuthCookies(w, access, refresh, rememberMe)

// 		json.NewEncoder(w).Encode(map[string]any{
// 			"status":  true,
// 			"message": "refreshed",
// 		})
// 	}
// }

// func (h handlerHTTP) HandlerLogout() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// ClearAuthCookies(w)
// 		log.Println("🚪 Logout initiated") // Debug log
// 		utils.ClearAuthCookies(w)         // Penting!

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]any{
// 			"status":  true,
// 			"message": "logged out",
// 		})
// 	}
// }

// func (h *handlerHTTP) HandlerAuthMe() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// 1. Ambil token dari cookie
// 		cookie, err := r.Cookie("access_token")
// 		if err != nil || cookie.Value == "" {
// 			http.Error(w, "unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// 2. Verifikasi JWT
// 		claims, err := utils.ValidateJWT(cookie.Value) // sesuaikan dengan fungsi JWT Anda
// 		if err != nil {
// 			http.Error(w, "invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// 3. Ambil user dari database
// 		user, err := h.Auth.GetUserByID(claims.UserID)
// 		if err != nil {
// 			http.Error(w, "user not found", http.StatusInternalServerError)
// 			return
// 		}
// 		userLoggedin := UserLoggedIn{
// 			ID:              user.ID,
// 			Username:        user.Username,
// 			Email:           user.Email,
// 			Role:            user.Role,
// 			SekolahTenantID: user.SekolahTenantID,
// 		}
// 		// 4. Ambil sekolah (opsional)
// 		var sekolahModel *models.SekolahTenant
// 		var SekolahTenant SekolahTenant
// 		if user.Role != "superadmin" {
// 			sekolahModel, err = h.repoSekolah.GetSekolahByTenantId(user.SekolahTenantID)
// 			if err != nil {
// 				http.Error(w, "Gagal ambil data sekolah", http.StatusUnauthorized)
// 				return
// 			}

// 			SekolahTenant.NamaSekolah = sekolahModel.NamaSekolah
// 			SekolahTenant.Npsn = sekolahModel.NPSN
// 			SekolahTenant.EnkripID = sekolahModel.EnkripID
// 		}

// 		// 5. Kirim response JSON
// 		w.Header().Set("Content-Type", "application/json")
// 		_ = json.NewEncoder(w).Encode(LoginResponse{
// 			Status:        true,
// 			Message:       "Login berhasil",
// 			User:          userLoggedin,
// 			SekolahTenant: &SekolahTenant,
// 		})
// 	}
// }

// func (h *handlerHTTP) HandlerSekolah() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }
