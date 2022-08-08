package promo

import (
	"payment/models"
	"time"
)

type PromoFormatter struct {
	ID         int64     `json:"id"`
	Code       string    `json:"promo_code"`
	Discount   int64     `json:"discount_amount"`
	Counter    int64     `json:"promo_code_counter"`
	Limited    bool      `json:"is_limited"`
	Available  bool      `json:"available"`
	ExpiryDate time.Time `json:"date_expire"`
}

func FormatPromo(promo models.PromoCode) PromoFormatter {
	formatter := PromoFormatter{}
	formatter.ID = promo.ID
	formatter.Code = promo.Code
	formatter.Discount = promo.Discount
	formatter.Counter = promo.Counter
	formatter.Limited = promo.Limited
	formatter.Available = promo.Available
	formatter.ExpiryDate = promo.ExpiryDate
	return formatter
}
