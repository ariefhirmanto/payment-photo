package usecase

import (
	"fmt"
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
	}

	res, err := p.coreapi.ChargeTransaction(chargeReq)
	fmt.Printf("%+v\n", chargeReq)
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", err)
	if err != nil {
		return res, err
	}

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
