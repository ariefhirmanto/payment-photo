package promo

import (
	"time"
)

type InputPromoCodeRequest struct {
	Discount   int64     `json:"discount_amount" binding:"required"`
	Code       string    `json:"promo_code" binding:"required"`
	Limited    bool      `json:"is_limited"`
	Counter    int64     `json:"counter"`
	ExpiryDate time.Time `json:"date_expire" binding:"required"`
	Available  bool      `json:"available" binding:"required"`
}

type InputPromoCodeID struct {
	ID int64 `uri:"id" binding:"required"`
}

type InputPromoCodeByCode struct {
	Code string `uri:"code" binding:"required"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
