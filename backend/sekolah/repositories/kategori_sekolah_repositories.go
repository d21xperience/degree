package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewKategoriSekolahRepository(db *gorm.DB) *GenericRepository[models.KategoriSekolah] {
	return NewGenericRepository[models.KategoriSekolah](db, "tabel_kategori_sekolah")
}
