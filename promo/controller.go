package promo

import "github.com/gin-gonic/gin"

type Controller interface {
	CreatePromoCode(c *gin.Context)
	GetPromoCodeByID(c *gin.Context)
	GetPromoCodeByCode(c *gin.Context)
	ClaimPromo(c *gin.Context)
	DeletePromo(c *gin.Context)
}
