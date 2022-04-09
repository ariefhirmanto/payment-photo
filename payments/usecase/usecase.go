package usecase

import (
	"context"
	"payment/models"
	"strconv"

	"github.com/midtrans/midtrans-go"
	snap "github.com/midtrans/midtrans-go/snap"
)

var s snap.Client

type paymentUsecase struct {
	ClientKey string
	ServerKey string
	Env       string
}

func NewPaymentUsecase(clientKey string, serverKey string, env string) *paymentUsecase {
	return &paymentUsecase{
		ClientKey: clientKey,
		ServerKey: serverKey,
		Env:       env,
	}
}

func (u *paymentUsecase) InitializeSnapClient() {
	env := midtrans.Sandbox
	if u.Env == "production" {
		env = midtrans.Sandbox
	}

	s.New(u.ServerKey, env)
}

func (u *paymentUsecase) GetPaymentURL(transaction models.Transaction) (string, error) {
	s.Options.SetContext(context.Background())

	resp, err := s.CreateTransactionUrl(u.GenerateSnapReq(transaction))
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (u *paymentUsecase) GenerateSnapReq(transaction models.Transaction) *snap.Request {

	// Initiate Customer address
	// custAddress := &midtrans.CustomerAddress{
	// 	FName:       "John",
	// 	LName:       "Doe",
	// 	Phone:       "081234567890",
	// 	Address:     "Baker Street 97th",
	// 	City:        "Jakarta",
	// 	Postcode:    "16000",
	// 	CountryCode: "IDN",
	// }

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(transaction.ID)),
			GrossAmt: transaction.Amount,
		},
		// CreditCard: &snap.CreditCardDetails{
		// 	Secure: true,
		// },
		// CustomerDetail: &midtrans.CustomerDetails{
		// 	FName:    "John",
		// 	LName:    "Doe",
		// 	Email:    "john@doe.com",
		// 	Phone:    "081234567890",
		// },
		// EnabledPayments: snap.AllSnapPaymentType,
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
		},
		// Items: &[]midtrans.ItemDetails{
		// 	{
		// 		ID:    "ITEM1",
		// 		Price: 200000,
		// 		Qty:   1,
		// 		Name:  "Someitem",
		// 	},
		// },
	}
	return snapReq
}
