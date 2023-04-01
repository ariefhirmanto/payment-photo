package frame

import "payment/models"

type Repository interface {
	CreateFrameImage(input models.Frame) (models.Frame, error)
	GetByID(ID int64) (models.Frame, error)
	GetAll() ([]models.Frame, error)
	GetByCategoryID(input int64) ([]models.Frame, error)
	Delete(input models.Frame) (bool, error)
	Update(input models.Frame) (models.Frame, error)
	GetFrameByLocation(input string) ([]models.Frame, error)
}
