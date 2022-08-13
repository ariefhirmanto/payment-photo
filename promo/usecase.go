package promo

import "payment/models"

type Usecase interface {
	CreatePromoCode(input FormPromoCodeRequest) (models.PromoCode, error)
	FindByID(input InputPromoCodeID) (models.PromoCode, error)
	FindByPromoCode(input InputPromoCodeByCode) (models.PromoCode, error)
	ClaimPromo(input InputPromoCodeByCode) (models.PromoCode, error)
	DeletePromo(input InputPromoCodeID) error
	GetAllPromo() ([]models.PromoCode, error)
	UpdatePromoCode(inputID InputPromoCodeID, inputData FormUpdatePromoCodeRequest) (models.PromoCode, error)
	UpdatePromoActivation(inputID InputPromoCodeID, inputData FormPromoActivation) error
}
