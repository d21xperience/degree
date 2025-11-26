package services

import (
	"context"
	"log"
	"net/http"

	"auth_service/config"
	pb "auth_service/generated"
	"auth_service/http_handler"
	"auth_service/interceptor"
	"auth_service/jwt"
	"auth_service/models"
	"auth_service/queue"
	"auth_service/repositories"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServiceServer dengan Redis Client sebagai Dependency Injection
type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	// RedisClient *redis.Client // Tambahkan Redis sebagai field
	repoSekolah   repositories.SekolahRepository
	authService   AuthService
	repoProfile   repositories.GenericRepository[models.UserProfile]
	repoUser      repositories.UserRepository
	rQueue        queue.RedisEnqueue
	pvKey         any
	jwtManager    *jwt.Manager
	cookieHandler *http_handler.CookieHandler
}

func NewAuthServiceServer(pvKey any) *AuthServiceServer {
	repoAuth := repositories.NewUserRepository(config.DB)
	authService := NewAuthService(repoAuth)
	repoSekolah := repositories.NewSekolahRepository(config.DB)
	repoProfile := repositories.NewUserProfileRepository(config.DB)
	repoUser := repositories.NewUserRepository(config.DB)
	rQueue := queue.NewRedisEnqueue(config.RDB)
	jwtManager := jwt.NewManager()
	conf := http_handler.CookieConfig{
		Secure:   true,
		Domain:   "localhost",
		SameSite: http.SameSiteLaxMode,
	}

	cookieHandler := http_handler.NewCookieHandler(conf)
	return &AuthServiceServer{
		authService:   authService,
		repoSekolah:   repoSekolah,
		repoProfile:   *repoProfile,
		repoUser:      repoUser,
		rQueue:        *rQueue,
		pvKey:         pvKey,
		jwtManager:    jwtManager,
		cookieHandler: cookieHandler,
	}
}

// type SchoolRegistration struct {
// 	SchoolName string `json:"school_name"`
// 	AdminEmail string `json:"admin_email"`
// }

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Validate credentials
	username := req.GetUsername()
	email := req.GetEmail()
	password := req.GetPassword()

	if password == "" || (username == "" && email == "") {
		return nil, status.Error(codes.InvalidArgument, "username/email dan password harus diisi")
	}

	var user *models.User
	var err error

	if email != "" {
		user, err = s.authService.Login(email, password)
	} else {
		user, err = s.authService.Login(username, password)
	}

	if err != nil {
		log.Printf("Login error: %v", err)
		return nil, status.Error(codes.Unauthenticated, "Username/email atau password salah")
	}

	// Ambil data Sekolah
	asalSekolah, err := s.repoSekolah.GetSekolahByTenantId(user.SekolahTenantID)
	if err != nil {
		log.Printf("Error fetching sekolahTenant: %v", err)
		return nil, status.Error(codes.NotFound, "Sekolah tidak ditemukan")
	}

	// Generate tokens
	// token, exp, err := utils.GenerateTokenRS256(s.pvKey, user, asalSekolah)
	// if err != nil {
	// 	log.Printf("token gen error: %v", err)
	// 	// return nil, fmt.Errorf("token gen error")
	// 	return nil, status.Error(codes.Internal, "token gen error")
	// }
	accessToken, refreshToken, err := s.jwtManager.GenerateTokens(user, asalSekolah)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate tokens")
	}

	return &pb.LoginResponse{
		AccessToken: accessToken,
		// ExpiresAt:    exp,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceServer) AuthMe(ctx context.Context, _ *pb.Empty) (*pb.UserProfile, error) {
	// ✅ Ambil user_id dari context (yang di-inject oleh interceptor)
	userID, ok := ctx.Value(interceptor.UserIDKey).(string)
	if !ok || userID == "" {
		return nil, status.Error(codes.Internal, "user_id not found in context")
	}

	// ✅ Ambil data user dari DB/repo
	user, err := s.repoUser.FindByID(userID)
	if err != nil {
		// if errors.Is(err, domain.ErrUserNotFound) {
		// 	return nil, status.Error(codes.Unauthenticated, "user not found") // bisa jadi token palsu
		// }
		return nil, status.Error(codes.Internal, "failed to fetch user")
	}

	// ✅ Konversi ke proto
	return &pb.UserProfile{
		UserId: user.ID,
		// Username: user.Username,
		// Email:    user.Email,
	}, nil
}

// func (s *AuthServiceServer) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.LoginResponse, error) {

// 	refreshID := req.RefreshToken
// 	if refreshID == "" {
// 		return nil, status.Error(codes.InvalidArgument, "empty refresh token")
// 	}

// 	// cek refresh token di Redis
// 	userID, err := s.redis.Get(ctx, "refresh:"+refreshID).Result()
// 	if err == redis.Nil {
// 		return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
// 	} else if err != nil {
// 		return nil, status.Error(codes.Internal, "redis error")
// 	}

// 	// generate access token baru
// 	accessToken, kid, err := GenerateAccessToken(userID, s.privateKey)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "cannot generate access token")
// 	}

// 	// rotate refresh token
// 	newRefreshID := uuid.New().String()
// 	s.redis.Set(ctx, "refresh:"+newRefreshID, userID, 30*24*time.Hour)
// 	s.redis.Del(ctx, "refresh:"+refreshID)

// 	return &authpb.RefreshTokenResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: newRefreshID,
// 	}, nil
// }

// func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Sekolah", "User"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Ambil data dari request
// 	user := req.GetUser()
// 	sekolah := req.GetSekolah()
// 	// Cek apakah sekolah sudah ada
// 	var sekolahModel *models.SekolahTenant
// 	sekolahModel, err = s.repoSekolah.GetSekolahByNPSN(sekolah.Npsn)
// 	if err != nil {
// 		if errors.Is(err, repositories.ErrRecordNotFound) {
// 			if user.Role == "admin" {
// 				// Buat sekolah baru
// 				sekolahModel = &models.SekolahTenant{
// 					NamaSekolah:   sekolah.NamaSekolah,
// 					NPSN:          sekolah.Npsn,
// 					EnkripID:      *sekolah.EnkripId,
// 					Kecamatan:     sekolah.Kecamatan,
// 					Kabupaten:     sekolah.Kabupaten,
// 					Propinsi:      sekolah.Propinsi,
// 					KodeKecamatan: sekolah.KodeKecamatan,
// 					KodeKab:       sekolah.KodeKab,
// 					KodeProp:      sekolah.KodeProp,
// 					AlamatJalan:   sekolah.AlamatJalan,
// 				}
// 				// err = s.repoSekolah.CreateSekolah(sekolahModel)

// 				// if err != nil {
// 				// 	log.Printf("Gagal membuat sekolah: %v", err)
// 				// 	return nil, fmt.Errorf("gagal membuat sekolah: %w", err)
// 				// }

// 			} else {
// 				// Pendaftaran siswa
// 				return nil, fmt.Errorf("sekolah belum terdaftar di aplikasi")
// 			}
// 		} else {
// 			log.Printf("Server error saat mencari sekolah: %v", err)
// 			return nil, fmt.Errorf("server error: %w", err)
// 		}
// 	}
// 	// Hubungkan user dengan sekolah
// 	userModel := &models.User{
// 		Username: utils.GenerateUsername(user.Email, sekolah.Npsn),
// 		Email:    user.Email,
// 		Role:     user.Role,
// 		// SekolahTenantID: sekolahModel.ID,
// 		Password: *user.Password,
// 	}

// 	// Cek jika role user adalah admin dan apakah sudah ada admin
// 	switch userModel.Role {
// 	case "admin":
// 		adminExists, err := s.authService.IsAdminExists(sekolahModel.ID)
// 		if err != nil {
// 			log.Printf("Error mengecek admin: %v", err)
// 			return nil, fmt.Errorf("server error: %w", err)
// 		}
// 		if adminExists {
// 			return nil, fmt.Errorf("admin sudah ada untuk sekolah ini")
// 		}

// 		// Registrasi admin
// 		// if err := s.authService.RegisterAdmin(userModel); err != nil {
// 		// 	log.Printf("Error registrasi admin: %v", err)
// 		// 	return nil, fmt.Errorf("gagal registrasi admin: %w", err)
// 		// }
// 		// cari username
// 		emailExists, err1 := s.repoUser.EmailExists(user.Email)
// 		if err1 != nil {
// 			return nil, fmt.Errorf("terjadi error: %s", err)
// 		}
// 		if emailExists {
// 			return nil, fmt.Errorf("alamat email sudah digunakan oleh sekolah lain")
// 		}

// 		if !emailExists {
// 			// Jika email tidak ada lanjutkan pendaftaran
// 			// Mendaftarkan sekolah
// 			err = s.repoSekolah.CreateSekolah(sekolahModel)
// 			if err != nil {
// 				log.Printf("Gagal membuat sekolah: %v", err)
// 				return nil, fmt.Errorf("gagal membuat sekolah: %w", err)
// 			}
// 			// mendaftarkan username
// 			userModel.SekolahTenantID = sekolahModel.ID
// 			err = s.authService.RegisterAdmin(userModel)
// 			if err != nil {
// 				log.Printf("Error registrasi admin: %v", err)
// 				return nil, fmt.Errorf("gagal registrasi admin: %w", err)
// 			}

// 		}
// 		if err := s.rQueue.EnqueueInitSekolahTask(*sekolahModel); err != nil {
// 			log.Printf("Gagal enqueue task: %v", err)
// 			return nil, fmt.Errorf("gagal enqueue initSekolahService: %w", err)
// 		}
// 		if err := s.rQueue.EnqueueInitSCTask(*sekolahModel, userModel.ID); err != nil {
// 			log.Printf("Gagal enqueue task: %v", err)
// 			return nil, fmt.Errorf("gagal enqueue initSCService: %w", err)
// 		}

// 	case "siswa":
// 		// Registrasi siswa
// 		if err := s.authService.Register(userModel); err != nil {
// 			log.Printf("Error registrasi siswa: %v", err)
// 			return nil, fmt.Errorf("gagal registrasi siswa: %w", err)
// 		}
// 	}

// 	// Hubungkan user dengan profil
// 	userProfileModel := &models.UserProfile{
// 		UserID: userModel.ID,
// 	}

// 	if err := s.repoProfile.Save(ctx, userProfileModel, "public"); err != nil {
// 		log.Printf("Error membuat user profile: %v", err)
// 		return nil, fmt.Errorf("server error saat membuat user profile")
// 	}

// 	var response *pb.RegisterResponse
// 	if userModel.Role == "admin" {
// 		token, err := utils.GenerateJWT(userModel, 25*time.Minute)
// 		if err != nil {
// 			return nil, errors.New("failed to generate token")
// 		}
// 		response = &pb.RegisterResponse{
// 			Token: token,
// 			Ok:    true,
// 			User: &pb.User{
// 				Id:       userModel.ID,
// 				Username: userModel.Username,
// 				// Email:     userModel.Email,
// 				Role:            userModel.Role,
// 				SekolahTenantId: userModel.SekolahTenantID,
// 			},
// 			SekolahTenant: &pb.SekolahTenant{
// 				NamaSekolah: sekolahModel.NamaSekolah,
// 				// EnkripId: &sekolahModel.EnkripID,
// 			},
// 		}
// 	} else {
// 		response = &pb.RegisterResponse{
// 			Ok: true,
// 		}
// 	}
// 	log.Println("User berhasil didaftarkan")
// 	return response, nil
// }

// Digunakan untuk membuat user baru dengan role siswa
// func (s *AuthServiceServer) CreateUsers(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Users"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	users := req.GetUsers()
// 	for _, v := range users {
// 		username := utils.GenerateUsername(v.Sekolah.NamaSekolah, v.UserProfile.Nama)
// 		pass, err := utils.GeneratePassword(4)
// 		if err != nil {
// 			return nil, err
// 		}
// 		newUser := models.User{
// 			Username:        username,
// 			Email:           v.Email,
// 			Password:        pass,
// 			SekolahTenantID: v.User.SekolahTenantId,
// 			Role:            v.User.Role,
// 		}
// 		err = s.authService.Register(&newUser)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// Hubungkan dengan user profile
// 		userId := newUser.ID
// 		userProfile := models.UserProfile{
// 			UserID:   userId,
// 			Nama:     v.UserProfile.Nama,
// 			JK:       v.UserProfile.Jk,
// 			Phone:    &v.UserProfile.Phone,
// 			TptLahir: &v.UserProfile.TptLahir,
// 			// TglLahir:  v.UserProfile.TglLahir,
// 			AlamatJalan: &v.UserProfile.AlamatJalan,
// 			KotaKab:     &v.UserProfile.KotaKab,
// 		}
// 		err = s.repoProfile.Save(ctx, &userProfile, "public")
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return &pb.CreateUsersResponse{
// 		Message: "ok",
// 	}, nil
// }
// func (s *AuthServiceServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
// 	return &pb.ResetPasswordResponse{
// 		Message: "password berhasil direset",
// 	}, nil
// }
// func (s *AuthServiceServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SekolahId"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var roleSiswa bool
// 	if req.Role == "" {
// 		roleSiswa = true
// 	}

// 	var usersModel []models.User
// 	if roleSiswa {
// 		usersModel, err = s.repoUser.GetUsers("siswa", req.GetSekolahTenantId())
// 		if err != nil {
// 			return nil, err
// 		}

// 	} // tambahkan else jika ingin menampilkan role admin
// 	res := utils.ConvertModelsToPB(usersModel, func(model models.User) *pb.User {
// 		return &pb.User{
// 			Username:  model.Username,
// 			Password:  model.InitialPassword,
// 			LastLogin: proto.String(utils.TimeToString(*model.LastLogin, "2006-01=12")),
// 		}
// 	})

// 	return &pb.GetUsersResponse{
// 		Users: res,
// 	}, nil

// }

// func (s *AuthServiceServer) GetSekolah(ctx context.Context, req *pb.GetSekolahRequest) (*pb.GetSekolahResponse, error) {
// 	sekolah, err := s.repoSekolah.GetSekolahByNPSN(req.GetNpsn())
// 	if err != nil {
// 		log.Printf("Error fetching sekolah: %v", err)
// 		if errors.Is(err, repositories.ErrRecordNotFound) {
// 			// return nil, status.Error(codes.NotFound, "school not found")
// 			return &pb.GetSekolahResponse{
// 				Status:  false,
// 				Message: "Sekolah tidak ditemukan",
// 			}, nil

// 		}
// 		// return nil, status.Error(codes.Internal, "internal error while fetching data")
// 		return &pb.GetSekolahResponse{
// 			Status:  false,
// 			Message: "Terjadi kesalah pada database",
// 		}, status.Error(codes.Internal, "internal error while fetching data")
// 	}
// 	return &pb.GetSekolahResponse{
// 		Status:  true,
// 		Message: "Sekolah sudah terdaftar",
// 		Nama:    sekolah.NamaSekolah,
// 		SekolahData: &pb.SekolahTenant{
// 			Id:            &sekolah.ID,
// 			EnkripId:      &sekolah.EnkripID,
// 			Kecamatan:     sekolah.Kecamatan,
// 			Kabupaten:     sekolah.Kabupaten,
// 			Propinsi:      sekolah.Propinsi,
// 			KodeKecamatan: sekolah.Kecamatan,
// 			AlamatJalan:   sekolah.AlamatJalan,
// 			KodeKab:       sekolah.KodeKab,
// 			KodeProp:      sekolah.KodeProp,
// 			NamaSekolah:   sekolah.NamaSekolah,
// 			Npsn:          sekolah.NPSN,
// 		},
// 	}, nil
// }
