package frame

import "payment/models"

type Usecase interface {
	SaveFrameImage(input FormInputFrame, fileLocation string) (models.Frame, error)
	GetFrameByID(input InputFrameID) (models.Frame, error)
	GetFrameByCategoryID(input InputCategoryID) ([]models.Frame, error)
	DeleteFrame(input InputFrameID) error
	GetAllFrame() ([]models.Frame, error)
	ChangeStatusFrame(input InputFrameID) error
	GetFrameByCategoryName(input InputCategoryName) ([]models.Frame, error)
	GetFrameByLocation(input InputLocationName) ([]models.Frame, error)
	GetFrameByName(input InputFrameName) (models.Frame, error)
}
