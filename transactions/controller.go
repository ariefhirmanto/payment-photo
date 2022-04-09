package transactions

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateTransaction(c *gin.Context)
	GetTransactionByID(c *gin.Context)
	GetNotification(c *gin.Context)
}
