package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewContractDataRepository(db *gorm.DB) *GenericRepository[models.Contract] {
	return NewGenericRepository[models.Contract](db, "contracts")
}
