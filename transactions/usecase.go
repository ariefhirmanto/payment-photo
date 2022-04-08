package transactions

import "payment/models"

type Usecase interface {
	CreateTransaction(input InputTransactionRequest) (models.Transaction, error)
	FindByID(input InputTransactionID) (models.Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}
