package services

import (
	"auth_service/models"
	"auth_service/repositories"
	"auth_service/utils"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type AuthService interface {
	IsAdminExists(schoolTenantID uint32) (bool, error)
	Register(user *models.User) error
	RegisterAdmin(user *models.User) error
	Login(username, password string) (*models.User, error)
	GetUserByID(userId int64) (*models.User, error)
	// GenerateToken(userID int, role string) (string, error)
	// SetAuthCookies(w http.ResponseWriter, accessToken, refreshToken string)
	// ClearAuthCookies(w http.ResponseWriter)
}

// AuthServiceImpl is the implementation of AuthService
type authServiceImpl struct {
	repo repositories.UserRepository
	// secretKey any
}

func NewAuthService(as repositories.UserRepository) AuthService {
	return &authServiceImpl{repo: as}
}

// IsAdminExists cek apakah admin sudah adah ada pada sekolah
func (s *authServiceImpl) IsAdminExists(schoolTenantID uint32) (bool, error) {
	admin, err := s.repo.FindUserByRoleAndSchoolID("admin", schoolTenantID)
	if err != nil {
		// Return false if no admin found or error is not nil
		if err == repositories.ErrUserNotFound {
			return false, nil
		}
		return false, err
	}
	return admin != nil, nil
}

func (s *authServiceImpl) Register(user *models.User) error {
	// Cek apakah username sudah ada
	var lanjutkan bool
	existingUser, err := s.repo.FindByUsername(user.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lanjutkan = true
		} else {
			// Tangani error jika terjadi kesalahan dalam mencari user
			return fmt.Errorf("failed to check existing username: %w", err)
		}
	}

	if existingUser != nil {
		return errors.New("username already exists")
	}
	// Simpan user baru
	if !lanjutkan {
		return fmt.Errorf("failed to save user: %w", err)
	}
	// user.Password, _ = utils.EncryptPassword(user.Password) // Encrypt password
	// return s.repositories.Save(user)
	// Enkripsi password
	user.InitialPassword = &user.Password
	encryptedPasswordChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	go func() {
		encryptedPassword, err := utils.EncryptPassword(user.Password)
		if err != nil {
			errorChan <- err
			return
		}
		encryptedPasswordChan <- encryptedPassword
	}()

	// Tunggu enkripsi selesai
	select {
	case user.Password = <-encryptedPasswordChan:
		// Simpan admin baru
		return s.repo.Save(user)
	case err = <-errorChan:
		return err
	}

}
func (s *authServiceImpl) RegisterAdmin(user *models.User) error {
	// Cek apakah email sudah ada dengan query efisien
	emailExists, err := s.repo.EmailExists(user.Email) // Hanya cek keberadaan email
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email already exists")
	}

	// Enkripsi password
	encryptedPasswordChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	go func() {
		encryptedPassword, err := utils.EncryptPassword(user.Password)
		if err != nil {
			errorChan <- err
			return
		}
		encryptedPasswordChan <- encryptedPassword
	}()

	// Tunggu enkripsi selesai
	select {
	case user.Password = <-encryptedPasswordChan:
		// Simpan admin baru
		return s.repo.Save(user)
	case err = <-errorChan:
		return err
	}
}

func (s *authServiceImpl) Login(identifier, password string) (*models.User, error) {
	var user *models.User
	var err error

	// Cek apakah identifier adalah email
	if utils.IsEmail(identifier) {
		user, err = s.repo.FindByEmail(identifier)
	} else {
		user, err = s.repo.FindByUsername(identifier)
	}

	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	err = s.repo.UpdateLastLogin(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	return user, nil
} 

func (s *authServiceImpl) GetUserByID(userId int64) (*models.User, error) {
	cekUser, err := s.repo.FindByID(strconv.FormatInt(userId, 10))
	if err != nil {
		return nil, err
	}
	return cekUser, nil
}
 