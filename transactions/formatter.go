package transactions

import "payment/models"

type TransactionFormatter struct {
	ID        int64  `json:"id"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
	QRCodeURL string `json:"qr_code"`
	UUID      string `json:"trx_id"`
}

func FormatTransaction(transaction models.Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.QRCodeURL = transaction.QRCodeURL
	formatter.UUID = transaction.TrxId
	return formatter
}
