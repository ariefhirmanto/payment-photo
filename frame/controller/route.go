package controller

import (
	"payment/frame"
	category "payment/frame_category"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc frame.Usecase, categoryUC category.Usecase, env string) {
	h := NewFrameController(uc, categoryUC, env)

	transactionEndpoints := router.Group("/api/v1/frame")

	{
		transactionEndpoints.GET("/", h.GetAllFrame)
		transactionEndpoints.GET("/category/:category_name", h.GetFrameByCategoryName)
		transactionEndpoints.GET("/location/:location", h.GetFrameByLocation)
		transactionEndpoints.GET("/frame/:id", h.GetFrameByID)
	}
}
