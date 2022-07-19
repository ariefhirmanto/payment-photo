package controller

import (
	"net/http"
	"payment/auth"
	"payment/helper"
	"payment/users"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUC      users.Usecase
	authService auth.Service
}

func NewUserController(usersUC users.Usecase, authService auth.Service) *userController {
	return &userController{userUC: usersUC, authService: authService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Register account failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userUC.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse(
			"Register account failed",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse(
			"Register account failed",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := users.FormatUser(newUser, token)

	response := helper.APIResponse(
		"Account has been registered",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var input users.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Login failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedInUser, err := h.userUC.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(
			"Login failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse(
			"Login failed",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(loggedInUser, token)
	response := helper.APIResponse(
		"Login Success",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
