package controller

import (
	"payment/payments"
)

type paymentController struct {
	paymentUC payments.Usecase
}

func NewPaymentController(usecase payments.Usecase) *paymentController {
	return &paymentController{paymentUC: usecase}
}

// func (p *paymentController)
