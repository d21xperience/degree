package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewKategoriSekolahLogRepository(db *gorm.DB) *GenericRepository[models.KategoriSekolahLog] {
	return NewGenericRepository[models.KategoriSekolahLog](db, "kategori_sekolah_log")
}
