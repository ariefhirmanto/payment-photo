package frame

import (
	"fmt"
	"payment/models"
)

type FrameFormatter struct {
	ID       int64           `json:"frame_id"`
	Name     string          `json:"name"`
	Category models.Category `json:"category"`
	Location string          `json:"location"`
	Url      string          `json:"url"`
	Counter  int             `json:"counter"`
}

func FormatFrame(frame models.Frame) FrameFormatter {
	fmt.Printf("Frame in formatter: %+v\n", frame)
	formatter := FrameFormatter{}
	formatter.ID = frame.ID
	formatter.Name = frame.Name
	formatter.Category = frame.Category
	formatter.Location = frame.Location
	formatter.Url = frame.Url
	formatter.Counter = frame.Counter
	fmt.Printf("formatter: %+v\n", formatter)
	return formatter
}

func FormatFrames(frames []models.Frame) []FrameFormatter {
	if len(frames) == 0 {
		return []FrameFormatter{}
	}

	var frameFormatter []FrameFormatter

	for _, frame := range frames {
		formatter := FormatFrame(frame)
		frameFormatter = append(frameFormatter, formatter)
	}

	return frameFormatter
}
