package transactions

import "payment/models"

type Repository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetByID(ID int64) (models.Transaction, error)
}
