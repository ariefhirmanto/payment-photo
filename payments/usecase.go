package payments

import (
	"payment/models"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Usecase interface {
	// InitializeSnapClient()
	GetPaymentURL(models.PaymentTransaction) (string, error)
	GetQRCode(models.PaymentTransaction) (*coreapi.ChargeResponse, error)
	// GenerateSnapReq(transaction models.Transaction) *snap.Request
}
