package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"sekolah/models"

	"gorm.io/gorm"
)

type SekolahRepository interface {
	Save(ctx context.Context, sekolah *models.Sekolah, schemaName string) error
	Find(ctx context.Context, schemaName string) (*models.Sekolah, error)
	FindByID(ctx context.Context, sekolahID string, schemaName string) (*models.Sekolah, error)
	Update(ctx context.Context, sekolah *models.Sekolah, schemaName string) error
	Delete(ctx context.Context, sekolahID string, schemaName string) error
	GetKategoriSekolah(ctx context.Context, namaSekolah string) (*models.BentukPendidikan, error)
	FindWithRelations(
		ctx context.Context,
		schemaName string,
		joins []string,
		preloads []string,
		exactConditions map[string]interface{},
		customConditions []struct {
			Query string
			Args  []interface{}
		},
		groupByColumns []string,
		orderBy []string,
	) ([]models.Sekolah, error)
}

type sekolahRepositoryImpl struct {
	// schemaRepository SchemaRepository
	db *gorm.DB
}

// NewSekolahRepository membuat instance baru dari SekolahRepository
func NewSekolahRepository(dB *gorm.DB) SekolahRepository {
	return &sekolahRepositoryImpl{
		db: dB,
		// schemaRepository: NewSchemaRepository(dB),
	}
}

var namaTabel = "tabel_sekolah"

func (r *sekolahRepositoryImpl) Save(ctx context.Context, sekolah *models.Sekolah, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}
		//  Gunakan `tx.Table(schemaName + ".sekolahs")` agar GORM tahu schema yang benar
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).Create(sekolah).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *sekolahRepositoryImpl) Find(ctx context.Context, schemaName string) (*models.Sekolah, error) {
	var sekolah models.Sekolah

	//  Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	//  Gunakan `tx.Table(schemaName + ".tabel_sekolah")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
		First(&sekolah).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &sekolah, nil
}
func (r *sekolahRepositoryImpl) FindByID(ctx context.Context, sekolahID string, schemaName string) (*models.Sekolah, error) {
	var sekolah models.Sekolah

	//  Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	//  Gunakan `tx.Table(schemaName + ".tabel_sekolah")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
		First(&sekolah, "sekolah_id = ?", sekolahID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &sekolah, nil
}

// Update (Memperbarui Data Sekolah)
func (r *sekolahRepositoryImpl) Update(ctx context.Context, sekolah *models.Sekolah, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		//  Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
			Where("sekolah_id = ?", sekolah.SekolahID).
			Updates(sekolah).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data Sekolah berdasarkan ID)
func (r *sekolahRepositoryImpl) Delete(ctx context.Context, sekolahID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		//  Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
			Where("sekolah_id = ?", sekolahID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
func (r *sekolahRepositoryImpl) GetKategoriSekolah(ctx context.Context, namaSekolah string) (*models.BentukPendidikan, error) {
	var kategori models.BentukPendidikan

	err := r.db.WithContext(ctx).
		Model(&models.BentukPendidikan{}).
		Where("? ILIKE '%' || nama || '%'", namaSekolah).
		Where("aktif = ?", true).
		Order("LENGTH(nama) DESC").
		Limit(1).
		Take(&kategori).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Kembalikan kategori default "Lainnya" dengan pointer struct
		return &models.BentukPendidikan{Nama: "Lainnya"}, nil
	} else if err != nil {
		return nil, err
	}

	return &kategori, nil
}

func (r *sekolahRepositoryImpl) FindWithRelations(
	ctx context.Context,
	schemaName string,
	joins []string,
	preloads []string,
	exactConditions map[string]interface{},
	customConditions []struct {
		Query string
		Args  []interface{}
	},
	groupByColumns []string,
	orderBy []string,
) ([]models.Sekolah, error) {
	var results []models.Sekolah
	tx := r.db.WithContext(ctx)

	// Set schema ke schemaName
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// Hapus tx.Distinct() → tidak digunakan

	// Join relasi jika ada
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Preload relasi jika ada
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// GROUP BY jika ada
	if len(groupByColumns) > 0 {
		tx = tx.Group(strings.Join(groupByColumns, ", "))
	}

	// ORDER BY jika ada
	if len(orderBy) > 0 {
		tx = tx.Order(strings.Join(orderBy, ", "))
	}

	// Kondisi exact
	if len(exactConditions) > 0 {
		tx = tx.Where(exactConditions)
	}

	// Kondisi custom
	for _, cond := range customConditions {
		tx = tx.Where(cond.Query, cond.Args...)
	}

	// Eksekusi query
	if err := tx.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
