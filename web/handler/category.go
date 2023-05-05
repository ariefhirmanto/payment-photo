package handler

import (
	"fmt"
	"net/http"
	category "payment/frame_category"
	user "payment/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryUC category.Usecase
	userUC     user.Usecase
}

func NewCategoryHandler(categoryUC category.Usecase, userUC user.Usecase) *categoryHandler {
	return &categoryHandler{categoryUC: categoryUC, userUC: userUC}
}

func (h *categoryHandler) Index(c *gin.Context) {
	categories, err := h.categoryUC.GetAllCategory()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "category_index.html", gin.H{"categories": categories})
}

func (h *categoryHandler) New(c *gin.Context) {
	input := category.FormInputCategory{}
	c.HTML(http.StatusOK, "category_new.html", input)
}

func (h *categoryHandler) Create(c *gin.Context) {
	var input category.FormInputCategory

	err := c.ShouldBind(&input)
	if err != nil {
		fmt.Printf("Error bind: %+v", err)
		c.HTML(http.StatusOK, "category_new.html", input)
		return
	}

	categoryInput := category.FormInputCategory{}
	categoryInput.Name = input.Name
	categoryInput.InterRowPadding = input.InterRowPadding
	categoryInput.InterColPadding = input.InterColPadding
	categoryInput.TopFramePadding = input.TopFramePadding
	categoryInput.CustomPadding = input.CustomPadding
	categoryInput.ImageID = input.ImageID
	categoryInput.Width = input.Width
	categoryInput.Height = input.Height
	categoryInput.IsColumnMirrored = input.IsColumnMirrored
	categoryInput.IsNoCut = input.IsNoCut
	categoryInput.IsSeasonal = input.IsSeasonal

	_, err = h.categoryUC.CreateCategory(categoryInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/category")
}

func (h *categoryHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingCategory, err := h.categoryUC.FindByID(category.InputCategoryID{ID: int64(id)})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := category.FormUpdateCategory{}
	input.ID = existingCategory.ID
	input.Name = existingCategory.Name
	input.InterRowPadding = existingCategory.InterRowPadding
	input.InterColPadding = existingCategory.InterColPadding
	input.TopFramePadding = existingCategory.TopFramePadding
	input.CustomPadding = existingCategory.CustomPadding
	input.ImageID = existingCategory.ImageID
	input.Width = existingCategory.Width
	input.Height = existingCategory.Height
	input.IsColumnMirrored = existingCategory.IsColumnMirrored
	input.IsNoCut = existingCategory.IsNoCut
	input.IsSeasonal = existingCategory.IsSeasonal

	c.HTML(http.StatusOK, "category_edit.html", input)
}

func (h *categoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input category.FormUpdateCategory

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	updateInput := category.FormUpdateCategory{}
	updateInput.ID = (int64(id))
	updateInput.Name = input.Name
	updateInput.InterRowPadding = input.InterRowPadding
	updateInput.InterColPadding = input.InterColPadding
	updateInput.TopFramePadding = input.TopFramePadding
	updateInput.CustomPadding = input.CustomPadding
	updateInput.ImageID = input.ImageID
	updateInput.Width = input.Width
	updateInput.Height = input.Height
	updateInput.IsColumnMirrored = input.IsColumnMirrored
	updateInput.IsNoCut = input.IsNoCut
	updateInput.IsSeasonal = input.IsSeasonal

	_, err = h.categoryUC.UpdateCategory(updateInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/category")
}

func (h *categoryHandler) DeleteRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	input := category.InputCategoryID{}
	input.ID = int64(id)
	c.HTML(http.StatusOK, "category_index.html", input)
}

func (h *categoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input category.FormUpdateCategory

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	err = h.categoryUC.DeleteCategory(category.InputCategoryID{ID: int64(id)})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/category")
}
