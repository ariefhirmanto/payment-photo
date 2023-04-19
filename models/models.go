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
	Location    string
}

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PromoCode struct {
	ID         int64
	Code       string
	Discount   int64
	Counter    int64
	Limited    bool
	Available  bool
	Duration   int
	ExpiryDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Frame struct {
	ID         int64
	Category   Category
	CategoryID int64
	Url        string
	Name       string
	Location   string
	Available  bool
	Counter    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Category struct {
	ID               int64
	Name             string
	InterRowPadding  int64
	TopFramePadding  int64
	InterColPadding  int64
	CustomPadding    int64
	ImageID          int64
	Width            int64
	Height           int64
	IsColumnMirrored bool
	IsNoCut          bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Location struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
