package repositories

import (
	"auth_service/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(userID string) (*models.User, error)
	FindUserByRoleAndSchoolID(role string, schoolID int32) (*models.User, error)
	Save(user *models.User) error
	EmailExists(email string) (bool, error)
	UpdateLastLogin(userID int64) error
	GetUsers(role string, schoolID uint32) ([]models.User, error)
	WithTx(tx *gorm.DB) UserRepository
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) WithTx(tx *gorm.DB) UserRepository {
	return &userRepository{db: tx}
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("ID = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Save(user *models.User) error {
	return r.db.Create(user).Error
}

// FindUserByRoleAndSchoolID fetches a user with the given role and school ID
func (r *userRepository) FindUserByRoleAndSchoolID(role string, schoolID int32) (*models.User, error) {
	var user models.User
	err := r.db.Where("role = ? AND sekolah_tenant_id = ?", role, schoolID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) EmailExists(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) UpdateLastLogin(userID int64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("last_login", time.Now()).Error
}

func (r *userRepository) GetUsers(role string, schoolID uint32) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("role = ? AND sekolah_id = ?", role, schoolID).Find(&users).Error
	if err != nil {
		return nil, err
	}

	// Jika tidak ada data ditemukan, kembalikan error khusus
	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return users, nil
}
