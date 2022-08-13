package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"payment/helper"
	"payment/promo"
)

type promoController struct {
	promoUC promo.Usecase
}

func NewPromoController(promoUC promo.Usecase) *promoController {
	return &promoController{promoUC: promoUC}
}

func (t *promoController) CreatePromoCode(c *gin.Context) {
	var input promo.FormPromoCodeRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Create promo code failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// input.ID = uuid.New().String()
	newPromo, err := t.promoUC.CreatePromoCode(input)
	if err != nil {
		response := helper.APIResponse(
			"Create promo code failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := promo.FormatPromo(newPromo)
	response := helper.APIResponse(
		"Success create promo code",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (t *promoController) GetPromoCodeByID(c *gin.Context) {
	var input promo.InputPromoCodeID
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Get promo failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	promoCode, err := t.promoUC.FindByID(input)
	if err != nil {
		response := helper.APIResponse(
			"Get promo failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := promo.FormatPromo(promoCode)
	response := helper.APIResponse(
		"Success get data promo code",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (t *promoController) GetPromoCodeByCode(c *gin.Context) {
	var input promo.InputPromoCodeByCode
	err := c.ShouldBindUri(&input)
	if err != nil {
		log.Printf("%+v", err)
		response := helper.APIResponse(
			"Get promo code by code failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	promoCode, err := t.promoUC.FindByPromoCode(input)
	if err != nil {
		response := helper.APIResponse(
			"Get promo code by code failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := promo.FormatPromo(promoCode)
	response := helper.APIResponse(
		"Success get data promo by code",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (t *promoController) ClaimPromo(c *gin.Context) {
	var input promo.InputPromoCodeByCode
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to claim promo",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	promoCode, err := t.promoUC.ClaimPromo(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to claim promo",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := promo.FormatPromo(promoCode)
	response := helper.APIResponse(
		"Success claimed promo",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (t *promoController) DeletePromo(c *gin.Context) {
	var input promo.InputPromoCodeID

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to delete promo code",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = t.promoUC.DeletePromo(input)
	if err != nil {
		errors := helper.FormatValidationError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to delete promo code",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
