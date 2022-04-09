package usecase

import (
	"payment/models"
	"payment/transactions"
	"strconv"
)

type transactionUsecase struct {
	TransactionRepo transactions.Repository
}

func NewTransactionUsecase(repo transactions.Repository) *transactionUsecase {
	return &transactionUsecase{TransactionRepo: repo}
}

func (u *transactionUsecase) CreateTransaction(input transactions.InputTransactionRequest) (models.Transaction, error) {
	transaction := models.Transaction{}
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.Code = ""

	newTransaction, err := u.TransactionRepo.CreateTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}

	// paymentTransaction := models.PaymentTransaction{
	// 	ID:     newTransaction.ID,
	// 	Amount: newTransaction.Amount,
	// }

	// paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction)
	// if err != nil {
	// 	return newTransaction, err
	// }

	// newTransaction.PaymentURL = paymentURL
	newTransaction, err = u.TransactionRepo.UpdateTransaction(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (u *transactionUsecase) FindByID(input transactions.InputTransactionID) (models.Transaction, error) {
	transaction, err := u.TransactionRepo.GetByID(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (u *transactionUsecase) ProcessPayment(input transactions.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := u.TransactionRepo.GetByID(int64(transaction_id))
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "captured" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	_, err = u.TransactionRepo.UpdateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}
