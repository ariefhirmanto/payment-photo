package usecase

import "gorm.io/gorm"

type locationRepository struct {
	locationDB *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *locationRepository {
	return &locationRepository{
		locationDB: db,
	}
}
