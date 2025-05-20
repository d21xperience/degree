package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewInfoIjazahRepository(db *gorm.DB) *GenericRepository[models.TabelInformasiIjazah] {
	return NewGenericRepository[models.TabelInformasiIjazah](db, "tabel_informasi_ijazah")
}
