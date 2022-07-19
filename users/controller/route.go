package controller

import (
	"payment/auth"
	"payment/users"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc users.Usecase, auth auth.Service) {
	h := NewUserController(uc, auth)

	transactionEndpoints := router.Group("/api/v1/admin")

	{
		transactionEndpoints.POST("/", h.RegisterUser)
		transactionEndpoints.POST("/login", h.Login)
	}
}
