package frame_category

type InputCategoryRequest struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	InterRowPadding int64  `json:"inter_row_padding"`
	TopFramePadding int64  `json:"top_frame_padding"`
	InterColPadding int64  `json:"inter_col_padding"`
	CustomPadding   int64  `json:"custom_padding"`
}

type FormInputCategory struct {
	ID              int64  `form:"id"`
	Name            string `form:"name"`
	InterRowPadding int64  `form:"inter_row_padding"`
	TopFramePadding int64  `form:"top_frame_padding"`
	InterColPadding int64  `form:"inter_col_padding"`
	CustomPadding   int64  `form:"custom_padding"`
	ImageID         int64  `form:"image_id"`
	Width           int64  `form:"width"`
	Height          int64  `form:"height"`
	Error           error
}

type FormUpdateCategory struct {
	ID              int64  `form:"id"`
	Name            string `form:"name"`
	InterRowPadding int64  `form:"inter_row_padding"`
	TopFramePadding int64  `form:"top_frame_padding"`
	InterColPadding int64  `form:"inter_col_padding"`
	CustomPadding   int64  `form:"custom_padding"`
	ImageID         int64  `form:"image_id"`
	Width           int64  `form:"width"`
	Height          int64  `form:"height"`
	Error           error
}

type InputCategoryID struct {
	ID int64 `uri:"id" binding:"required"`
}

type InputCategoryName struct {
	Name string `uri:"name" binding:"required"`
}
