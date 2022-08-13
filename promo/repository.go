package promo

import "payment/models"

type Repository interface {
	CreatePromoCode(promo models.PromoCode) (models.PromoCode, error)
	UpdatePromoCode(promo models.PromoCode) (models.PromoCode, error)
	GetByID(ID int64) (models.PromoCode, error)
	GetByPromoCode(code string) (models.PromoCode, error)
	GetAll() ([]models.PromoCode, error)
	Update(input models.PromoCode) (models.PromoCode, error)
}
