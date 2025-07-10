package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewBCPlatformRepository(db *gorm.DB) *GenericRepository[models.BCPlatform] {
	return NewGenericRepository[models.BCPlatform](db, "ref.blockchain_platform")
}
