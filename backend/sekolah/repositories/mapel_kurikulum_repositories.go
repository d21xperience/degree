package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewMapelKurikulumRepository(db *gorm.DB) *GenericRepository[models.MataPelajaranKurikulum] {
	return NewGenericRepository[models.MataPelajaranKurikulum](db, "mata_pelajaran_kurikulum")
}
