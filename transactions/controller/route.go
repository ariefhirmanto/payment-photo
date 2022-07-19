package controller

import (
	"payment/transactions"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc transactions.Usecase) {
	h := NewTransactionControllers(uc)

	transactionEndpoints := router.Group("/api/v1/transaction")

	{
		transactionEndpoints.POST("/", h.CreateTransaction)
		transactionEndpoints.GET("/:id", h.GetTransactionByID)
		transactionEndpoints.GET("/trx/:trx_id", h.GetTransactionByTrxID)
		transactionEndpoints.POST("/notification", h.GetNotification)
	}
}
