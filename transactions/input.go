package transactions

import "errors"

type InputTransactionRequest struct {
	ID     int64  `json:"id" binding:"required"`
	Amount int64  `json:"amount" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type InputTransactionID struct {
	ID int64 `uri:"id" binding:"required"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

func (input InputTransactionRequest) ValidateInput() error {
	if input.Amount == 0 {
		return errors.New("amount can't be 0")
	}

	return nil
}

func (input InputTransactionID) ValidateInput() error {
	if input.ID == 0 {
		return errors.New("ID can't be 0")
	}

	return nil
}
