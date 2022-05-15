package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"payment/helper"
	"payment/transactions"
)

type transactionController struct {
	transactionUC transactions.Usecase
}

func NewTransactionControllers(transactionUC transactions.Usecase) *transactionController {
	return &transactionController{transactionUC: transactionUC}
}

func (t *transactionController) CreateTransaction(c *gin.Context) {
	var input transactions.InputTransactionRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Create transaction failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// input.ID = uuid.New().String()
	newTransaction, err := t.transactionUC.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse(
			"Create transaction failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := transactions.FormatTransaction(newTransaction)
	response := helper.APIResponse(
		"Success create transaction",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (t *transactionController) GetTransactionByID(c *gin.Context) {
	var input transactions.InputTransactionID
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Get transaction failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := t.transactionUC.FindByID(input)
	if err != nil {
		response := helper.APIResponse(
			"Get transaction failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := transactions.FormatTransaction(transaction)
	response := helper.APIResponse(
		"Success get data transaction",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (t *transactionController) GetNotification(c *gin.Context) {
	var input transactions.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to process notification",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = t.transactionUC.ProcessPayment(input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to process notification",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
