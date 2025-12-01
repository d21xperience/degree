package repositories

import (
	"auth_service/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// type SekolahTenantRepository interface {
// 	CreateSekolahTenant(*models.SekolahTenant) error
// 	// GetSekolahTenant(query SekolahTenantQuery) (*models.SekolahTenant, error)
// 	GetSekolahTenantByNPSN(npsn string) (*models.SekolahTenant, error)
// 	GetSekolahTenantByTenantId(tenantId int32) (*models.SekolahTenant, error)

// 	// WithTx mengembalikan instance repositori yang menggunakan transaksi
// 	WithTx(tx *gorm.DB) SekolahTenantRepositoryImpl
// }

// type SekolahTenantRepositoryImpl struct {
// 	DB *gorm.DB
// }

type SekolahTenantQuery struct {
	Npsn            string
	SekolahTenantID int
}

// func NewSekolahTenantRepository(db *gorm.DB) SekolahTenantRepository {
// 	return &SekolahTenantRepositoryImpl{DB: db}
// }

// func (r *SekolahTenantRepositoryImpl) WithTx(tx *gorm.DB) SekolahTenantRepositoryImpl {
// 	return &SekolahTenantRepository{DB: tx}
// }

// func (r *SekolahTenantRepositoryImpl) CreateSekolahTenant(s *models.SekolahTenant) error {
// 	result := r.DB.Create(&s)
// 	if result.Error != nil {
// 		// Penangann Error jika terjadi duplikate
// 		return result.Error
// 	}
// 	return nil
// }

// func (r *SekolahTenantRepositoryImpl) GetSekolahTenantByNPSN(npsn string) (*models.SekolahTenant, error) {
// 	// Validasi: Pastikan minimal satu parameter ada
// 	var SekolahTenant models.SekolahTenant
// 	dbQuery := r.DB
// 	dbQuery = dbQuery.Where("npsn = ?", npsn)

// 	// Eksekusi query
// 	err := dbQuery.First(&SekolahTenant).Error
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, ErrRecordNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &SekolahTenant, nil
// }

// pkg/repositories/sekolah_repository.go

type SekolahTenantRepository interface {
	GetSekolahTenantByNPSN(npsn string) (*models.SekolahTenant, error)
	CreateSekolahTenant(sekolah *models.SekolahTenant) error

	GetSekolahTenantByTenantId(tenantId int32) (*models.SekolahTenant, error)
	// WithTx mengembalikan instance repositori yang menggunakan transaksi
	WithTx(tx *gorm.DB) SekolahTenantRepository
}

type sekolahRepository struct {
	db *gorm.DB
}

func NewSekolahTenantRepository(db *gorm.DB) SekolahTenantRepository {
	return &sekolahRepository{db: db}
}

func (r *sekolahRepository) WithTx(tx *gorm.DB) SekolahTenantRepository {
	return &sekolahRepository{db: tx}
}

func (r *sekolahRepository) GetSekolahTenantByNPSN(npsn string) (*models.SekolahTenant, error) {
	var sekolah models.SekolahTenant
	if err := r.db.Where("npsn = ?", npsn).First(&sekolah).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &sekolah, nil
}

func (r *sekolahRepository) CreateSekolahTenant(sekolah *models.SekolahTenant) error {
	return r.db.Create(sekolah).Error
}

func (r *sekolahRepository) GetSekolahTenantByTenantId(tenantId int32) (*models.SekolahTenant, error) {
	// Validasi: Pastikan minimal satu parameter ada
	var SekolahTenant models.SekolahTenant
	dbQuery := r.db
	dbQuery = dbQuery.Where("id = ?", tenantId)

	// Eksekusi query
	err := dbQuery.First(&SekolahTenant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	return &SekolahTenant, nil
}

func (r *sekolahRepository) GetSekolahTenant(query SekolahTenantQuery) (*models.SekolahTenant, error) {
	// Validasi: Pastikan minimal satu parameter ada
	if query.Npsn == "" && query.SekolahTenantID == 0 {
		return nil, fmt.Errorf("minimal salah satu parameter (npsn atau SekolahTenant_id) harus disediakan")
	}

	var SekolahTenant models.SekolahTenant

	// Gunakan kedua parameter jika keduanya ada
	dbQuery := r.db
	if query.Npsn != "" && query.SekolahTenantID != 0 {
		dbQuery = dbQuery.Where("npsn = ? AND id = ?", query.Npsn, query.SekolahTenantID)
	} else {
		// Gunakan salah satu parameter jika hanya salah satu yang ada
		if query.Npsn != "" {
			dbQuery = dbQuery.Where("npsn = ?", query.Npsn)
		}
		if query.SekolahTenantID != 0 {
			dbQuery = dbQuery.Where("id = ?", query.SekolahTenantID)
		}
	}

	// Eksekusi query
	err := dbQuery.First(&SekolahTenant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	return &SekolahTenant, nil
}
