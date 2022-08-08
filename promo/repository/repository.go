package repository

import (
	"payment/models"

	"gorm.io/gorm"
)

type promoRepository struct {
	promoDB *gorm.DB
}

func NewPromoRepository(db *gorm.DB) *promoRepository {
	return &promoRepository{promoDB: db}
}

func (r *promoRepository) CreatePromoCode(promo models.PromoCode) (models.PromoCode, error) {
	err := r.promoDB.Create(&promo).Error
	if err != nil {
		return promo, err
	}

	return promo, nil
}

func (r *promoRepository) UpdatePromoCode(promo models.PromoCode) (models.PromoCode, error) {
	err := r.promoDB.Save(&promo).Error
	if err != nil {
		return promo, err
	}
	return promo, nil
}

func (r *promoRepository) GetByID(ID int64) (models.PromoCode, error) {
	var promoCode models.PromoCode
	// with index
	// err := r.promoDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&promoCode).Error
	// without indexing
	err := r.promoDB.Where("id = ?", ID).Find(&promoCode).Error
	if err != nil {
		return promoCode, err
	}

	return promoCode, nil
}

func (r *promoRepository) GetByPromoCode(code string) (models.PromoCode, error) {
	var promoCode models.PromoCode
	// with index
	// err := r.promoDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&promoCode).Error
	// without indexing
	err := r.promoDB.Where("code = ?", code).Find(&promoCode).Error
	if err != nil {
		return promoCode, err
	}

	return promoCode, nil
}
