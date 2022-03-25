package models

import "time"

type Transaction struct {
	ID          int64
	Amount      int64
	Status      string
	Code        string
	PaymentType int
	PaymentURL  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentTransaction struct {
	ID          int64
	Amount      int64
	PaymentType string
}
