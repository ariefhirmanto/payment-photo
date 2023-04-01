package frame_category

import "payment/models"

type Repository interface {
	CreateCategory(input models.Category) (models.Category, error)
	Update(input models.Category) (models.Category, error)
	GetByID(ID int64) (models.Category, error)
	GetAll() ([]models.Category, error)
	Delete(input models.Category) (bool, error)
	GetByName(input string) (models.Category, error)
}
