package controller

import (
	"payment/promo"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc promo.Usecase) {
	h := NewPromoController(uc)

	transactionEndpoints := router.Group("/api/v1/promo")

	{
		transactionEndpoints.GET("/:id", h.GetPromoCodeByID)
		transactionEndpoints.GET("/code/:code", h.GetPromoCodeByCode)
		transactionEndpoints.PUT("/claim/:code", h.ClaimPromo)
		transactionEndpoints.PUT("/:id", h.DeletePromo)
	}
}
