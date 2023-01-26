package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"payment/auth"
	"payment/helper"
	"payment/users"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func authMiddleware(authService auth.Service, userUsecase users.Usecase, productSwitcher bool) gin.HandlerFunc {
	if !productSwitcher {
		return func(c *gin.Context) {
			authHeader := c.GetHeader("Authorization")
			if !strings.Contains(authHeader, "Bearer") {
				response := helper.APIResponse("Unauthorized",
					http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			tokenString := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}
			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				response := helper.APIResponse("Unauthorized",
					http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response := helper.APIResponse("Unauthorized",
					http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			userID := int(claim["user_id"].(float64))
			user, err := userUsecase.GetUserByID(userID)
			if err != nil {
				response := helper.APIResponse("Unauthorized",
					http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			c.Set("currentUser", user)
		}
	}

	log.Printf("%+v", productSwitcher)

	return func(c *gin.Context) {
		c.Next()
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")
		fmt.Printf("%+v\n", userIDSession)

		if userIDSession == nil {
			fmt.Printf("%+v\n", userIDSession)
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}
