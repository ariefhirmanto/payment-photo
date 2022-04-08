package payments

import (
	"payment/models"

	"github.com/midtrans/midtrans-go/snap"
)

type Usecase interface {
	InitializeSnapClient()
	GetPaymentURL(transaction models.Transaction) (string, error)
	GenerateSnapReq(transaction models.Transaction) *snap.Request
}
