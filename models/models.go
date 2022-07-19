package models

import "time"

type Transaction struct {
	ID          int64
	Amount      int64
	Status      string
	PaymentType int
	QRCodeURL   string
	TrxId       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentTransaction struct {
	ID          int64
	Amount      int64
	PaymentType int
	TrxID       string
}

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
