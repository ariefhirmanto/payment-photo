package frame_category

import "payment/models"

type CategoryFormatter struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	InterRowPadding int64  `json:"inter_row_padding"`
	TopFramePadding int64  `json:"top_frame_padding"`
	InterColPadding int64  `json:"inter_col_padding"`
	CustomPadding   int64  `json:"custom_padding"`
}

func FormatCategory(category models.Category) CategoryFormatter {
	formatter := CategoryFormatter{}
	formatter.ID = category.ID
	formatter.Name = category.Name
	formatter.InterRowPadding = category.InterRowPadding
	formatter.TopFramePadding = category.TopFramePadding
	formatter.InterColPadding = category.InterColPadding
	formatter.CustomPadding = category.CustomPadding
	return formatter
}
