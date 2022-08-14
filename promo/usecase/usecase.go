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

func (u *promoUsecase) CreatePromoCode(input promo.FormPromoCodeRequest) (models.PromoCode, error) {
	promoCode := models.PromoCode{}
	promoCode.Code = input.Code
	promoCode.Discount = input.Discount

	if input.Discount <= 0 || input.Counter <= 0 || input.Duration <= 0 {
		log.Printf("[Promo][Usecase][CreatePromoCode] Unable to proceed negative input number for duration/discount/counter")
		return models.PromoCode{}, errors.New("input number can't be negative")
	}

	promoCode.Limited = input.Limited
	if !input.Limited {
		promoCode.Counter = -1
	} else {
		promoCode.Counter = input.Counter
	}

	promoCode.Duration = input.Duration
	promoCode.Available = input.Available
	if input.Available {
		promoCode.ExpiryDate = time.Now().AddDate(0, 0, input.Duration)
	} else {
		promoCode.ExpiryDate = time.Now()
	}

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

func (u *promoUsecase) GetAllPromo() ([]models.PromoCode, error) {
	promos, err := u.PromoRepo.GetAll()
	if err != nil {
		return promos, err
	}

	return promos, nil
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

	if !promo.Available {
		log.Printf("[Promo][Usecase][DeletePromo] Promo already deleted %+v", input.ID)
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

func (u *promoUsecase) UpdatePromoActivation(inputID promo.InputPromoCodeID, inputData promo.FormPromoActivation) error {
	promo, err := u.PromoRepo.GetByID(inputID.ID)
	if err != nil {
		log.Printf("[Promo][Usecase][DeletePromo] Error get promo transaction with ID %+v", inputID.ID)
		return err
	}

	if promo.Available == inputData.Available {
		log.Printf("[Promo][Usecase][DeletePromo] Promo with ID: %+v, can't change status: no change", inputID.ID)
		return err
	}

	promo.Available = inputData.Available
	if inputData.Available {
		promo.ExpiryDate = time.Now().AddDate(0, 0, promo.Duration)
	} else {
		promo.ExpiryDate = time.Now()
	}

	newPromo, err := u.PromoRepo.UpdatePromoCode(promo)
	if err != nil {
		log.Printf("[Promo][Usecase][DeletePromo] Failed to update promo with ID %+v", inputID.ID)
		return err
	}

	log.Printf("[Promo][Usecase][DeletePromo] Success delete promo %+v", newPromo)
	return nil
}

func (u *promoUsecase) UpdatePromoCode(inputID promo.InputPromoCodeID, inputData promo.FormUpdatePromoCodeRequest) (models.PromoCode, error) {
	promoCode, err := u.PromoRepo.GetByID(inputID.ID)
	if err != nil {
		return promoCode, err
	}

	promoCode.Code = inputData.Code
	promoCode.Discount = inputData.Discount
	promoCode.Counter = inputData.Counter
	promoCode.Limited = inputData.Limited
	promoCode.Duration = inputData.Duration

	promoCode.Available = inputData.Available
	if inputData.Available {
		promoCode.ExpiryDate = time.Now().AddDate(0, 0, inputData.Duration)
	} else {
		promoCode.ExpiryDate = time.Now()
	}

	updatedPromoCode, err := u.PromoRepo.Update(promoCode)
	if err != nil {
		return updatedPromoCode, err
	}

	return updatedPromoCode, nil
}
