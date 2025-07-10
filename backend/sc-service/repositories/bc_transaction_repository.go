package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewContractBCTransactionRepository(db *gorm.DB) *GenericRepository[models.BCTransaction] {
	return NewGenericRepository[models.BCTransaction](db, "transaksi_blockchain")
}
