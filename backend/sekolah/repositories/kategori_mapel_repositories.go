package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewKategoriMapelRepository(db *gorm.DB) *GenericRepository[models.KategoriMapel] {
	return NewGenericRepository[models.KategoriMapel](db, "tabel_kategori_mapel")
}
