package handler

import (
	"fmt"
	"net/http"
	"payment/helper"
	promo "payment/promo"
	user "payment/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type promoHandler struct {
	promoUC promo.Usecase
	userUC  user.Usecase
}

func NewPromoHandler(promoUC promo.Usecase, userUC user.Usecase) *promoHandler {
	return &promoHandler{promoUC: promoUC, userUC: userUC}
}

func (h *promoHandler) Index(c *gin.Context) {
	promos, err := h.promoUC.GetAllPromo()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "promo_index.html", gin.H{"promos": promos})
}

func (h *promoHandler) New(c *gin.Context) {
	input := promo.FormPromoCodeRequest{}
	c.HTML(http.StatusOK, "promo_new.html", input)
}

func (h *promoHandler) Create(c *gin.Context) {
	var input promo.FormPromoCodeRequest

	err := c.ShouldBind(&input)
	if err != nil {
		fmt.Printf("Error bind: %+v", err)
		c.HTML(http.StatusOK, "promo_new.html", input)
		return
	}

	promoCodeInput := promo.FormPromoCodeRequest{}
	promoCodeInput.Code = input.Code
	promoCodeInput.Discount = input.Discount
	promoCodeInput.Available = false
	promoCodeInput.Counter = input.Counter
	promoCodeInput.Limited = input.Limited
	promoCodeInput.Duration = input.Duration

	_, err = h.promoUC.CreatePromoCode(promoCodeInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/promo")
}

func (h *promoHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingPromo, err := h.promoUC.FindByID(promo.InputPromoCodeID{ID: int64(id)})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := promo.FormUpdatePromoCodeRequest{}
	input.ID = existingPromo.ID
	input.Code = existingPromo.Code
	input.Discount = existingPromo.Discount
	input.Counter = existingPromo.Counter
	input.Limited = existingPromo.Limited
	input.Available = existingPromo.Available
	input.Duration = 0
	// input.ExpiryDate = existingPromo.ExpiryDate

	c.HTML(http.StatusOK, "promo_edit.html", input)
}

func (h *promoHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input promo.FormUpdatePromoCodeRequest

	err := c.ShouldBind(&input)
	if err != nil {
		input.Code = helper.GenerateUUID()
		input.Discount = 0
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	updateInput := promo.FormUpdatePromoCodeRequest{}
	updateInput.ID = input.ID
	updateInput.Code = input.Code
	updateInput.Discount = input.Discount
	updateInput.Counter = input.Counter
	updateInput.Limited = input.Limited
	updateInput.Available = input.Available
	// updateInput.ExpiryDate = input.ExpiryDate

	_, err = h.promoUC.UpdatePromoCode(promo.InputPromoCodeID{ID: int64(id)}, updateInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/promo")
}

func (h *promoHandler) ActivatePage(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingPromo, err := h.promoUC.FindByID(promo.InputPromoCodeID{ID: int64(id)})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := promo.FormPromoActivation{}
	input.ID = existingPromo.ID
	input.Available = existingPromo.Available

	c.HTML(http.StatusOK, "promo_activation.html", input)
}

func (h *promoHandler) ActivationAction(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input promo.FormPromoActivation

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	err = h.promoUC.UpdatePromoActivation(promo.InputPromoCodeID{ID: int64(id)}, input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/promo")
}
