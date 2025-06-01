package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewKategoriKelasRepository(db *gorm.DB) *GenericRepository[models.KategoriKelas] {
	return NewGenericRepository[models.KategoriKelas](db, "tabel_kategori_kelas")
}
