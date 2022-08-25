package usecase

import (
	"log"
	"payment/helper"
	"payment/models"
	"payment/payments"
	"payment/transactions"
	"strconv"
)

type transactionUsecase struct {
	TransactionRepo transactions.Repository
	PaymentUC       payments.Usecase
}

func NewTransactionUsecase(repo transactions.Repository, payment payments.Usecase) *transactionUsecase {
	return &transactionUsecase{
		TransactionRepo: repo,
		PaymentUC:       payment,
	}
}

func (u *transactionUsecase) CreateTransaction(input transactions.InputTransactionRequest) (models.Transaction, error) {
	transaction := models.Transaction{}
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.PaymentType = input.PaymentType
	transaction.TrxId = helper.GenerateUUID()

	newTransaction, err := u.TransactionRepo.CreateTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := models.PaymentTransaction{
		ID:          newTransaction.ID,
		Amount:      newTransaction.Amount,
		PaymentType: newTransaction.PaymentType,
		TrxID:       newTransaction.TrxId,
		Location:    input.Location,
	}

	log.Printf("[Transactions][Usecase][CreateTransaction] Get payment transaction %+v", paymentTransaction)
	paymentURL, err := u.PaymentUC.GetQRCode(paymentTransaction)
	log.Printf("[Transactions][Usecase][CreateTransaction] Get payment transaction QR Code URL %+v", paymentURL)
	if err != nil {
		log.Printf("[Transactions][Usecase][CreateTransaction] Error get QRCode %+v", err)
		return newTransaction, err
	}

	newTransaction.QRCodeURL = ""
	if len(paymentURL.Actions) > 0 {
		newTransaction.QRCodeURL = paymentURL.Actions[0].URL
	}

	newTransaction, err = u.TransactionRepo.UpdateTransaction(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

// deprecated
func (u *transactionUsecase) FindByID(input transactions.InputTransactionID) (models.Transaction, error) {
	transaction, err := u.TransactionRepo.GetByID(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (u *transactionUsecase) CreateTransactionWithoutQRCode(input transactions.InputTransactionRequest) (models.Transaction, error) {
	transaction := models.Transaction{}
	transaction.Amount = input.Amount
	transaction.Status = "paid"
	transaction.PaymentType = 99
	transaction.TrxId = helper.GenerateUUID()
	log.Printf("[Transactions][Usecase][CreateTransactionWithoutQRCode] Bypass payment with input %+v", input)

	newTransaction, err := u.TransactionRepo.CreateTransaction(transaction)
	if err != nil {
		log.Printf("[Transactions][Usecase][CreateTransactionWithoutQRCode] Error get transactions with err %+v", err)
		return newTransaction, err
	}

	log.Printf("[Transactions][Usecase][CreateTransactionWithoutQRCode] Success create %+v", newTransaction)
	return newTransaction, nil
}

func (u *transactionUsecase) FindByTrxID(input transactions.InputTransactionTrxID) (models.Transaction, error) {
	transaction, err := u.TransactionRepo.GetByTrxID(input.TrxID)
	if err != nil {
		log.Printf("[Transactions][Usecase][FindByTrxID] Error get transactions with trx id %+v", input.TrxID)
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

func (u *transactionUsecase) ProcessPaymentV2(input transactions.TransactionNotificationInput) error {
	log.Printf("[Transactions][Usecase][ProcessPaymentV2] Get callback from Midtrans")
	transaction_id := input.OrderID

	transaction, err := u.TransactionRepo.GetByTrxID(transaction_id)
	if err != nil {
		log.Printf("[Transactions][Usecase][ProcessPaymentV2] Error get transactions with trx id %+v", input.OrderID)
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

	log.Printf("[Transactions][Usecase][ProcessPaymentV2] Success processed callback")
	return nil
}
