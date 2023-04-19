package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"payment/frame"
	category "payment/frame_category"
	"payment/helper"
	"payment/models"

	"github.com/gin-gonic/gin"
)

type frameController struct {
	frameUC    frame.Usecase
	categoryUC category.Usecase
	Env        string
}

func NewFrameController(frameUC frame.Usecase, categoryUC category.Usecase, Env string) *frameController {
	return &frameController{frameUC: frameUC, categoryUC: categoryUC, Env: Env}
}

func (h *frameController) UploadImage(c *gin.Context) {
	var input frame.FormInputFrame

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to upload frame",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	category, err := h.categoryUC.FindByIDForFrame(input.CategoryID)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	input.Category = category

	file, err := c.FormFile("file")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	parentDirectory := getDirectory(h.Env)
	path := fmt.Sprintf(parentDirectory+"images/%s/%s", category.Name, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.frameUC.SaveFrameImage(input, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(
		"Upload campaign image success",
		http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *frameController) GetAllFrame(c *gin.Context) {
	frames, err := h.frameUC.GetAllFrame()
	if err != nil {
		response := helper.APIResponse(
			"Get frame failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	for i := range frames {
		category, err := h.categoryUC.FindByIDForFrame(frames[i].CategoryID)
		if err != nil {
			response := helper.APIResponse(
				"Get frame failed",
				http.StatusUnprocessableEntity, "error", nil)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		frames[i].Category = category
	}

	formatter := frame.FormatFrames(frames)
	response := helper.APIResponse(
		"Success get data promo code",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *frameController) GetFrameByCategoryName(c *gin.Context) {
	var input frame.InputCategoryName
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Get frame by category name failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	frames, err := h.frameUC.GetFrameByCategoryName(input)
	if err != nil {
		response := helper.APIResponse(
			"Get frame failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	for i := range frames {
		category, err := h.categoryUC.FindByIDForFrame(frames[i].CategoryID)
		if err != nil {
			response := helper.APIResponse(
				"Get frame failed",
				http.StatusUnprocessableEntity, "error", nil)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		frames[i].Category = category
	}

	formatter := frame.FormatFrames(frames)
	response := helper.APIResponse(
		"Success get data frames by category",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *frameController) GetFrameByLocation(c *gin.Context) {
	var input frame.InputLocationName
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Get frame by category name failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	frames, err := h.frameUC.GetFrameByLocation(input)
	if err != nil {
		response := helper.APIResponse(
			"Get frame failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	for i := range frames {
		category, err := h.categoryUC.FindByIDForFrame(frames[i].CategoryID)
		if err != nil {
			response := helper.APIResponse(
				"Get frame failed",
				http.StatusUnprocessableEntity, "error", nil)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		frames[i].Category = category
	}

	formatter := frame.FormatFrames(frames)
	response := helper.APIResponse(
		"Success get data frames by location",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *frameController) GetFrameByID(c *gin.Context) {
	var input frame.InputFrameID
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Get frame by id failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.frameUC.GetFrameByID(input)
	if (err != nil || data == models.Frame{}) {
		response := helper.APIResponse(
			"Get frame failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	category, err := h.categoryUC.FindByIDForFrame(data.CategoryID)
	if err != nil {
		response := helper.APIResponse(
			"Get frame failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data.Category = category

	formatter := frame.FormatFrame(data)
	response := helper.APIResponse(
		"Success get data",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *frameController) GetFrameByName(c *gin.Context) {
	var input frame.InputFrameName
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Get frame by name failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.frameUC.GetFrameByName(input)
	if (err != nil || data == models.Frame{}) {
		response := helper.APIResponse(
			"Get frame failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	category, err := h.categoryUC.FindByIDForFrame(data.CategoryID)
	if err != nil {
		response := helper.APIResponse(
			"Get frame failed",
			http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data.Category = category

	formatter := frame.FormatFrame(data)
	response := helper.APIResponse(
		"Success get data",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func getDirectory(env string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	url := filepath.Dir(wd)
	if env != "local" {
		url = "/app" + filepath.Dir(wd)
	}

	return url
}
