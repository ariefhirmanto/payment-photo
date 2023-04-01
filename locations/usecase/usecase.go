package usecase

import (
	location "payment/locations"
)

type locationUsecase struct {
	LocationRepo location.Repository
}

func NewLocationUsecase(repo location.Repository) *locationUsecase {
	return &locationUsecase{
		LocationRepo: repo,
	}
}
