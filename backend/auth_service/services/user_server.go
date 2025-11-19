package services

import (
	"auth_service/config"
	pb "auth_service/generated"
	"auth_service/repositories"
	"context"
	"errors"
	"log"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo repositories.UserRepository
}

func NewUserUserServiceServer() *UserServer {
	repoUser := repositories.NewUserRepository(config.DB)
	return &UserServer{
		repo: repoUser,
	}
}

// GetUser - Mengambil pengguna berdasarkan UserID
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	userID := req.GetId()
	user, err := s.repo.FindByID(userID)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, errors.New("failed to retrieve user")
	}

	return &pb.User{
		Username: user.Username,
		Id:       user.ID,
		Role:     user.Role,
	}, nil
}

// func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UserProfilePhoto, error) {
// 	// Path ke file foto profil pengguna
// 	// filePath := fmt.Sprintf("photos/%d.jpg", req.GetUserId())

// 	// // Baca file foto
// 	// photoBytes, err := os.ReadFile(filePath)
// 	// if err != nil {
// 	// 	return nil, status.Errorf(codes.NotFound, "Photo not found")
// 	// }

// 	return &pb.ListUsersResponse{
// 		// SizeBytes:   photoBytes,
// 		// ContentType: "image/jpeg", // Ganti sesuai tipe file
// 	}, nil
// }

// // Create User - Memperbarui profil pengguna berdasarkan UserID
// func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfile, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received UpdateUserProfile request: %+v\n", req)

// 	// Cek apakah req atau req.UserProfile kosong
// 	if req == nil {
// 		log.Println("Request is nil")
// 		return nil, errors.New("invalid request: request is nil")
// 	}

// 	if req.UserProfile == nil {
// 		log.Println("UserProfile is nil in request")
// 		return nil, errors.New("invalid request: user profile is nil")
// 	}

// 	profile, err := s.repo.FindByID(ctx, utils.ConvertUintToString(uint64(req.UserId)), "public", "profile")
// 	if err != nil {
// 		log.Printf("Error fetching user profile: %v", err)
// 		return nil, errors.New("user profile not found")
// 	}
// 	// profi := req.GetProfile()
// 	// log.Println(profi)
// 	// Perbarui data profil berdasarkan input
// 	profile.Nama = req.UserProfile.Nama
// 	profile.JK = req.UserProfile.Jk
// 	// profile.Phone = req.UserProfile.Phone
// 	// profile.TptLahir = req.UserProfile.TptLahir
// 	// profile.AlamatJalan = req.UserProfile.AlamatJalan
// 	// profile.KotaKab = req.UserProfile.KotaKab
// 	// profile.Prov = req.UserProfile.Prov
// 	// profile.KodePos = req.UserProfile.KodePos
// 	// profile.NamaAyah = req.UserProfile.NamaAyah
// 	// profile.NamaIbu = req.UserProfile.NamaIbu

// 	// Simpan perubahan ke database
// 	userId := strconv.FormatInt(profile.UserID, 10)

// 	err = s.repo.Update(ctx, profile, "public", "user_id", userId)
// 	if err != nil {
// 		log.Printf("Error updating user profile: %v", err)
// 		return nil, errors.New("failed to update user profile")
// 	}

// 	return &pb.UserProfile{
// 		Message: "Updated",
// 	}, nil
// }
