package usecase

import (
	"log"
	"payment/models"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentMidtrans struct {
	coreapi    coreapi.Client
	snapClient snap.Client
}

func NewPaymentMidtrans(clientKey string, serverKey string, env string) *paymentMidtrans {
	cfgEnv := midtrans.Sandbox
	if env == "production" {
		cfgEnv = midtrans.Production
	}

	return &paymentMidtrans{
		coreapi: coreapi.Client{
			ClientKey: clientKey,
			ServerKey: serverKey,
			Env:       cfgEnv,
		},
		snapClient: snap.Client{
			ServerKey: serverKey,
			Env:       cfgEnv,
		},
	}
}

func (p *paymentMidtrans) GetQRCode(input models.PaymentTransaction) (*coreapi.ChargeResponse, error) {
	log.Printf("[Payments][Usecase][GetQRCode] Get transaction QR Code")
	paymentType := coreapi.PaymentTypeQris
	if input.PaymentType == 1 {
		paymentType = coreapi.PaymentTypeGopay
	}

	p.coreapi.New(p.coreapi.ServerKey, p.coreapi.Env)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: paymentType,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: input.TrxID,
			// GrossAmt: 30000,
			GrossAmt: input.Amount,
		},
		CustomField1: &input.Location,
	}
	log.Printf("[Payments][Usecase][GetQRCode] Request charge transaction QRCode %+v", chargeReq)

	res, err := p.coreapi.ChargeTransaction(chargeReq)
	log.Printf("[Payments][Usecase][GetQRCode] Get response charge transaction QR Code %+v", res)
	if err != nil {
		log.Printf("[Payments][Usecase][GetQRCode] Error get QRCode charge transaction with err %+v", err)
		return res, err
	}

	log.Printf("[Payments][Usecase][GetQRCode] Success create QR Code")
	return res, nil
}

func (p *paymentMidtrans) GetPaymentURL(input models.PaymentTransaction) (string, error) {
	paymentType := snap.PaymentTypeGopay
	p.snapClient.New(p.snapClient.ServerKey, p.snapClient.Env)

	chargeReq := &snap.Request{
		EnabledPayments: []snap.SnapPaymentType{
			paymentType,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(input.ID)),
			GrossAmt: 35000,
		},
	}

	res, err := p.snapClient.CreateTransactionUrl(chargeReq)
	if err != nil {
		return res, err
	}

	return res, nil
}
