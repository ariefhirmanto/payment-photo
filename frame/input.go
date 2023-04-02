package frame

import "payment/models"

type CreateFrameInput struct {
	CategoryID int64 `form:"category_id" binding:"required"`
}

type FormInputFrame struct {
	Categories []models.Category
	Category   models.Category
	Name       string `form:"name" binding:"required"`
	Location   string `form:"location" binding:"required"`
	CategoryID int64  `form:"category_id" binding:"required"`
	Counter    int    `form:"counter" binding:"required"`
	Error      error
}

type InputCategoryID struct {
	ID int64 `uri:"id" binding:"required"`
}

type InputCategoryName struct {
	CategoryName string `uri:"category_name" binding:"required"`
}

type InputLocationName struct {
	Location string `uri:"location" binding:"required"`
}

type InputFrameID struct {
	ID int64 `uri:"id" binding:"required"`
}

type InputFrameName struct {
	Name string `uri:"frame_name" binding:"required"`
}
