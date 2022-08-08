package usecase

import (
	"errors"
	"log"
	"payment/models"
	"payment/promo"
	"time"
)

type promoUsecase struct {
	PromoRepo promo.Repository
}

func NewPromoUsecase(repo promo.Repository) *promoUsecase {
	return &promoUsecase{
		PromoRepo: repo,
	}
}

func (u *promoUsecase) CreatePromoCode(input promo.InputPromoCodeRequest) (models.PromoCode, error) {
	promoCode := models.PromoCode{}
	promoCode.Code = input.Code
	promoCode.Discount = input.Discount
	promoCode.Limited = input.Limited
	if !input.Limited {
		promoCode.Counter = -1
	} else {
		promoCode.Counter = input.Counter
	}
	promoCode.ExpiryDate = input.ExpiryDate
	promoCode.Available = true

	promo, _ := u.PromoRepo.GetByPromoCode(input.Code)
	if (models.PromoCode{}) != promo {
		log.Printf("[Promo][Usecase][CreatePromoCode] Error promo code already exists %+v", input.Code)
		return promo, errors.New("promo code already exists")
	}

	savedPromoCode, err := u.PromoRepo.CreatePromoCode(promoCode)
	if err != nil {
		log.Printf("[Promo][Usecase][CreatePromoCode] Error creating transaction %+v", err)
		return savedPromoCode, err
	}

	log.Printf("[Promo][Usecase][CreatePromoCode] Success create transaction %+v", savedPromoCode)
	return savedPromoCode, nil
}

func (u *promoUsecase) FindByID(input promo.InputPromoCodeID) (models.PromoCode, error) {
	promo, err := u.PromoRepo.GetByID(input.ID)
	if err != nil {
		log.Printf("[Promo][Usecase][FindByID] Error get promo transaction with ID %+v", input.ID)
		return promo, err
	}

	return promo, nil
}

func (u *promoUsecase) FindByPromoCode(input promo.InputPromoCodeByCode) (models.PromoCode, error) {
	promo, err := u.PromoRepo.GetByPromoCode(input.Code)
	if err != nil {
		log.Printf("[Promo][Usecase][FindByPromoCode] Error get promo transaction with code %+v", input.Code)
		return promo, err
	}

	return promo, nil
}

func (u *promoUsecase) ClaimPromo(input promo.InputPromoCodeByCode) (models.PromoCode, error) {
	promo, err := u.PromoRepo.GetByPromoCode(input.Code)
	if err != nil {
		log.Printf("[Promo][Usecase][ClaimPromo] Error get promo transaction with code %+v", input.Code)
		return promo, err
	}

	if !promo.Available {
		log.Printf("[Promo][Usecase][ClaimPromo] Promo code not available with code %+v", input.Code)
		return promo, errors.New("Promo Code is not available")
	}

	if promo.ExpiryDate.Sub(time.Now()) < 0 {
		log.Printf("[Promo][Usecase][ClaimPromo] Promo already expired with code %+v", input.Code)
		return promo, errors.New("Promo Code already expired")
	}

	if promo.Limited {
		if promo.Counter == 0 {
			log.Printf("[Promo][Usecase][ClaimPromo] Promo code is out with code %+v", input.Code)
			return promo, errors.New("Promo Code is out")
		}
		promo.Counter -= 1
	}

	newPromo, err := u.PromoRepo.UpdatePromoCode(promo)
	if err != nil {
		log.Printf("[Promo][Usecase][ClaimPromo] Failed to update promo with code %+v", input.Code)
		return newPromo, err
	}

	log.Printf("[Promo][Usecase][ClaimPromo] Success claim promo %+v", newPromo)
	return promo, nil
}

func (u *promoUsecase) DeletePromo(input promo.InputPromoCodeID) error {
	promo, err := u.PromoRepo.GetByID(input.ID)
	if err != nil {
		log.Printf("[Promo][Usecase][DeletePromo] Error get promo transaction with ID %+v", input.ID)
		return err
	}

	promo.Available = false
	newPromo, err := u.PromoRepo.UpdatePromoCode(promo)
	if err != nil {
		log.Printf("[Promo][Usecase][DeletePromo] Failed to update promo with ID %+v", input.ID)
		return err
	}

	log.Printf("[Promo][Usecase][DeletePromo] Success delete promo %+v", newPromo)
	return nil
}
