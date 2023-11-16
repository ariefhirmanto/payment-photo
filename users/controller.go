package users

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Login(c *gin.Context)
	RegisterUser(c *gin.Context)
	ChangePassword(c *gin.Context)
}
