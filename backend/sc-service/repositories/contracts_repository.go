package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewContractDataRepository(db *gorm.DB) *GenericRepository[models.ContractData] {
	return NewGenericRepository[models.ContractData](db, "contracts")
}
