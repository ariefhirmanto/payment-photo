package transactions

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateTransaction(c *gin.Context)
	GetTransactionByID(c *gin.Context)
	GetNotification(c *gin.Context)
	GetNotificationV2(c *gin.Context)
	BypassNormalFlow(c *gin.Context)
	GetTransactionByTrxID(c *gin.Context)
}
