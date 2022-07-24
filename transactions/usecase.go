package transactions

import "payment/models"

type Usecase interface {
	CreateTransaction(input InputTransactionRequest) (models.Transaction, error)
	CreateTransactionWithoutQRCode(input InputTransactionRequest) (models.Transaction, error)
	FindByID(input InputTransactionID) (models.Transaction, error)
	FindByTrxID(input InputTransactionTrxID) (models.Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
	ProcessPaymentV2(input TransactionNotificationInput) error
}
