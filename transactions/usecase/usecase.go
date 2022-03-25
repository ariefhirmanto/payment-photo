package usecase

import (
	"payment/transactions"
)

type transactionUsecase struct {
	TransactionRepo transactions.Repository
}

func NewTransactionUsecase(repo transactions.Repository) *transactionUsecase {
	return &transactionUsecase{TransactionRepo: repo}
}
