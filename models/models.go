package models

import "time"

type Transaction struct {
	ID          int64
	Amount      int64
	Status      string
	PaymentType int
	QRCodeURL   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentTransaction struct {
	ID          int64
	Amount      int64
	PaymentType int
}
