package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	frame "payment/frame"
	category "payment/frame_category"
	user "payment/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type frameHandler struct {
	frameUC    frame.Usecase
	categoryUC category.Usecase
	userUC     user.Usecase
	env        string
}

func NewFrameHandler(frameUC frame.Usecase, userUC user.Usecase, categoryUC category.Usecase, env string) *frameHandler {
	return &frameHandler{frameUC: frameUC, userUC: userUC, categoryUC: categoryUC, env: env}
}

func (h *frameHandler) Index(c *gin.Context) {
	frames, err := h.frameUC.GetAllFrame()
	for i := range frames {
		category, err := h.categoryUC.FindByIDForFrame(frames[i].CategoryID)
		if err != nil {
			fmt.Printf("Error find by ID: %+v\n", err)
			c.HTML(http.StatusInternalServerError, "error.html", err)
			return
		}
		frames[i].Category = category
	}

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "frame_index.html", gin.H{"frames": frames})
}

func (h *frameHandler) New(c *gin.Context) {
	categories, err := h.categoryUC.GetAllCategory()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err)
		return
	}

	input := frame.FormInputFrame{}
	input.Categories = categories
	c.HTML(http.StatusOK, "frame_new.html", input)
}

func (h *frameHandler) Create(c *gin.Context) {
	var input frame.FormInputFrame

	err := c.ShouldBind(&input)
	if err != nil {
		categories, e := h.categoryUC.GetAllCategory()
		if e != nil {
			c.HTML(http.StatusInternalServerError, "error.html", err)
			return
		}

		input.Categories = categories
		input.Error = err
		fmt.Printf("Input: %+v\n", input)

		c.HTML(http.StatusOK, "category_new.html", input)
		return
	}

	frameInput := frame.FormInputFrame{}
	frameInput.Name = input.Name
	frameInput.CategoryID = input.CategoryID
	frameInput.Location = input.Location
	frameInput.Counter = input.Counter
	fmt.Printf("Category: %+v\n", frameInput)

	category, err := h.categoryUC.FindByIDForFrame(input.CategoryID)
	if err != nil {
		fmt.Printf("Error find by ID: %+v\n", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	frameInput.Category = category

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("Error get file: %+v\n", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	parentDirectory := getDirectory(h.env)
	path := fmt.Sprintf(parentDirectory+"images/%s/%s", category.Name, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		fmt.Printf("Error save: %+v\n", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	_, err = h.frameUC.SaveFrameImage(frameInput, path)
	if err != nil {
		fmt.Printf("Error save frame image: %+v\n", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/frame")
}

func (h *frameHandler) DeleteRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err)
		return
	}

	input := frame.InputFrameID{}
	input.ID = int64(id)
	c.HTML(http.StatusOK, "frame_index.html", input)
}

func (h *frameHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := h.frameUC.DeleteFrame(frame.InputFrameID{ID: int64(id)})
	if err != nil {
		fmt.Printf("Delete status frame: %+v\n", err)
		c.HTML(http.StatusInternalServerError, "error.html", err)
		return
	}

	log.Print("[Handler] Delete Frame success")
	c.Redirect(http.StatusFound, "/frame")
}

func (h *frameHandler) ChangeStatusRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err)
		return
	}

	input := frame.InputFrameID{}
	input.ID = int64(id)
	c.HTML(http.StatusOK, "frame_index.html", input)
}

func (h *frameHandler) ChangeStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", err)
		return
	}

	err = h.frameUC.ChangeStatusFrame(frame.InputFrameID{ID: int64(id)})
	if err != nil {
		fmt.Printf("Change status frame: %+v\n", err)
		c.HTML(http.StatusInternalServerError, "error.html", err)
		return
	}

	c.Redirect(http.StatusFound, "/frame")
}

func getDirectory(env string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	url := ""
	if env != "local" {
		url = "/app" + filepath.Dir(wd)
	}

	return url
}
