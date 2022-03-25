package controller

import (
	"payment/transactions"
)

type transactionController struct {
	transactionUC transactions.Usecase
}

func NewTransactionControllers(transactionUC transactions.Usecase) *transactionController {
	return &transactionController{transactionUC: transactionUC}
}
