package promo

import "payment/models"

type Usecase interface {
	CreatePromoCode(input InputPromoCodeRequest) (models.PromoCode, error)
	FindByID(input InputPromoCodeID) (models.PromoCode, error)
	FindByPromoCode(input InputPromoCodeByCode) (models.PromoCode, error)
	ClaimPromo(input InputPromoCodeByCode) (models.PromoCode, error)
	DeletePromo(input InputPromoCodeID) error
}
