package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"auth_service/config"
	pb "auth_service/generated"
	"auth_service/http_handler"
	"auth_service/jwt"
	"auth_service/models"
	"auth_service/queue"
	"auth_service/repositories"
	"auth_service/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// AuthServiceServer dengan Redis Client sebagai Dependency Injection
type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	// RedisClient *redis.Client // Tambahkan Redis sebagai field
	repoSekolah   repositories.SekolahTenantRepository
	authService   AuthService
	repoProfile   repositories.UserProfileRepository
	repoUser      repositories.UserRepository
	rQueue        queue.RedisEnqueue
	jwtManager    *jwt.Manager
	cookieHandler *http_handler.CookieHandler
	db            *gorm.DB
	// pvKey         any
}

func NewAuthServiceServer(jwtManager *jwt.Manager) *AuthServiceServer {
	db := config.DB
	repoAuth := repositories.NewUserRepository(db)
	authService := NewAuthService(repoAuth)
	repoSekolah := repositories.NewSekolahTenantRepository(db)
	repoProfile := repositories.NewUserProfileRepository(db)
	repoUser := repositories.NewUserRepository(db)
	rQueue := queue.NewRedisEnqueue(config.RDB)

	conf := http_handler.CookieConfig{
		Secure:   true,
		Domain:   "localhost",
		SameSite: http.SameSiteLaxMode,
	}

	cookieHandler := http_handler.NewCookieHandler(conf)
	return &AuthServiceServer{
		db:            db,
		authService:   authService,
		repoSekolah:   repoSekolah,
		repoProfile:   repoProfile,
		repoUser:      repoUser,
		rQueue:        *rQueue,
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

	accessToken, refreshToken, err := s.jwtManager.GenerateTokens(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate tokens")
	}

	return &pb.LoginResponse{
		AccessToken: accessToken,
		// ExpiresAt:    exp,
		RefreshToken: refreshToken,
		UserId:       user.ID,
		Email:        user.Email,
	}, nil
}

func (s *AuthServiceServer) AuthMe(ctx context.Context, _ *pb.Empty) (*pb.User, error) {
	// ✅ Ambil user_id dari context (yang di-inject oleh interceptor)
	userID, ok := ctx.Value(utils.CtxUserID).(string)
	// log.Printf("🔍 userID = %q, ok = %v", userID, ok) // <-- ini yang penting
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

	// Ambil data Sekolah
	asalSekolah, err := s.repoSekolah.GetSekolahTenantByTenantId(user.SekolahTenantID)
	if err != nil {
		log.Printf("Error fetching sekolahTenant: %v", err)
		return nil, status.Error(codes.NotFound, "Sekolah tidak ditemukan")
	}
	// ✅ Konversi ke proto
	return &pb.User{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		SekolahSlug: utils.NormalizeAnyString(asalSekolah.NamaSekolah),
		SekolahAsal: &pb.SekolahTenant{
			Id:          asalSekolah.ID,
			NamaSekolah: asalSekolah.NamaSekolah,
			Npsn:        asalSekolah.NPSN,
			EnkripId:    asalSekolah.EnkripID,
		},
	}, nil
}

// =========================================================
// REFAKTOR
// =========================================================
func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("[Register] Request received: %+v", req)

	// Validasi awal
	if req.GetUser() == nil || req.GetSekolah() == nil {
		return nil, status.Error(codes.InvalidArgument, "user and sekolah are required")
	}
	if req.GetUser().Role == "" && req.GetUser().Email == "" && req.GetUser().Password == "" {
		return nil, status.Error(codes.InvalidArgument, "role, email & password is required")
	}
	if req.GetSekolah().NamaSekolah == "" && req.GetSekolah().EnkripId == "" {
		return nil, status.Error(codes.InvalidArgument, "school name & enkrip id is required")
	}

	userReq := req.GetUser()
	sekolahReq := req.GetSekolah()

	// Mulai transaksi
	tx := s.db.Begin()
	if tx.Error != nil {
		log.Printf("[Register] Failed to begin transaction: %v", tx.Error)
		return nil, status.Error(codes.Internal, "database transaction failed")
	}

	// Auto rollback jika belum commit
	committed := false
	defer func() {
		if !committed {
			if err := tx.Rollback().Error; err != nil && !errors.Is(err, gorm.ErrInvalidTransaction) {
				log.Printf("[Register] Rollback failed (after error/panic): %v", err)
			}
		}
	}()

	// Dapatkan repositori & service dalam transaksi
	repoSekolahTx := s.repoSekolah.WithTx(tx)
	repoUserTx := s.repoUser.WithTx(tx)
	repoProfileTx := s.repoProfile.WithTx(tx)
	authServiceTx := s.authService.WithTx(tx)

	// 1. Cari atau buat sekolah (hanya admin boleh buat baru)
	var sekolahModel *models.SekolahTenant
	sekolahModel, err := repoSekolahTx.GetSekolahTenantByNPSN(sekolahReq.Npsn)
	if err != nil && !errors.Is(err, repositories.ErrRecordNotFound) {
		tx.Rollback()
		log.Printf("[Register] DB error checking school: %v", err)
		return nil, status.Error(codes.Internal, "failed to verify school")
	}

	if sekolahModel == nil {
		if userReq.Role != "admin" {
			tx.Rollback()
			return nil, status.Error(codes.PermissionDenied, "only admin can register new school")
		}

		sekolahModel = &models.SekolahTenant{
			NamaSekolah:   sekolahReq.NamaSekolah,
			NPSN:          sekolahReq.Npsn,
			EnkripID:      sekolahReq.EnkripId,
			Kecamatan:     sekolahReq.Kecamatan,
			Kabupaten:     sekolahReq.Kabupaten,
			Propinsi:      sekolahReq.Propinsi,
			KodeKecamatan: sekolahReq.KodeKecamatan,
			KodeKab:       sekolahReq.KodeKab,
			KodeProp:      sekolahReq.KodeProp,
			AlamatJalan:   sekolahReq.AlamatJalan,
		}

		if err := repoSekolahTx.CreateSekolahTenant(sekolahModel); err != nil {
			tx.Rollback()
			log.Printf("[Register] Failed to create school: %v", err)
			return nil, status.Error(codes.Internal, "failed to create school")
		}
		log.Printf("[Register] New school created (ID=%d, NPSN=%s)", sekolahModel.ID, sekolahModel.NPSN)
	}

	// 2. Cek duplikasi email (dalam transaksi → read consistency)
	emailExists, err := repoUserTx.EmailExists(userReq.Email)
	if err != nil {
		tx.Rollback()
		log.Printf("[Register] Email existence check failed: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	if emailExists {
		tx.Rollback()
		return nil, status.Error(codes.AlreadyExists, "email already registered")
	}

	// 3. Validasi khusus admin
	if userReq.Role == "admin" {
		adminExists, err := authServiceTx.IsAdminExists(sekolahModel.ID)
		if err != nil {
			tx.Rollback()
			log.Printf("[Register] Admin existence check failed: %v", err)
			return nil, status.Error(codes.Internal, "internal error")
		}
		if adminExists {
			tx.Rollback()
			return nil, status.Error(codes.AlreadyExists, "admin already exists for this school")
		}
	}

	// 4. Buat user
	userModel := &models.User{
		Username:        utils.GenerateUsername(userReq.Email, sekolahReq.Npsn),
		Email:           userReq.Email,
		Role:            userReq.Role,
		SekolahTenantID: sekolahModel.ID,
		Password:        userReq.Password, // akan di-hash di authService
	}

	var userID int64
	switch userReq.Role {
	case "admin":
		userID, err = authServiceTx.RegisterAdmin(userModel)
	// case "siswa":
	// 	userID, err = authServiceTx.RegisterSiswa(userModel)
	default:
		tx.Rollback()
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("unsupported role: %s", userReq.Role))
	}

	if err != nil {
		tx.Rollback()
		log.Printf("[Register] User registration failed: %v", err)
		return nil, status.Error(codes.Internal, "user registration failed")
	}
	userModel.ID = userID

	// 5. Buat profil
	userProfile := &models.UserProfile{
		UserID: userID,
	}
	if err := repoProfileTx.Save(userProfile); err != nil {
		tx.Rollback()
		log.Printf("[Register] Profile creation failed for user %d: %v", userID, err)
		return nil, status.Error(codes.Internal, "failed to create user profile")
	}

	// 6. ✅ Semua sukses → commit
	if err := tx.Commit().Error; err != nil {
		log.Printf("[Register] Transaction commit failed: %v", err)
		return nil, status.Error(codes.Internal, "transaction commit failed")
	}

	// // ✅ Tandai committed (opsional, untuk defer check)
	// tx.Statement.Vars["committed"] = true

	// 7. Enqueue background jobs (di luar transaksi — async)
	if userReq.Role == "admin" {
		if err := s.rQueue.EnqueueInitSekolahTask(*sekolahModel); err != nil {
			log.Printf("[WARN] Failed to enqueue initSekolahTask: %v", err)
		}
		if err := s.rQueue.EnqueueInitSCTask(*sekolahModel, userID); err != nil {
			log.Printf("[WARN] Failed to enqueue initSCService: %v", err)
		}
	}

	// 8. Bangun respons
	resp := &pb.RegisterResponse{
		Status:  true,
		Message: "Akun berhasil dibuat",
	}

	if userReq.Role == "admin" {
		// resp.User = &pb.User{
		// 	Id:              uint32(userID),
		// 	Username:        userModel.Username,
		// 	Role:            userModel.Role,
		// 	SekolahTenantId: sekolahModel.ID,
		// }
		// resp.SekolahTenant = &pb.SekolahTenant{
		// 	Id:          sekolahModel.ID,
		// 	NamaSekolah: sekolahModel.NamaSekolah,
		// 	Npsn:        sekolahModel.NPSN,
		// }
	}

	log.Printf("[Register] SUCCESS: user=%d (%s), school=%d", userID, userReq.Role, sekolahModel.ID)
	return resp, nil
}

// // Digunakan untuk membuat user baru dengan role siswa
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
