package frame_category

import "payment/models"

type Usecase interface {
	CreateCategory(input FormInputCategory) (models.Category, error)
	GetAllCategory() ([]models.Category, error)
	FindByID(input InputCategoryID) (models.Category, error)
	FindByIDForFrame(input int64) (models.Category, error)
	FindByName(input InputCategoryName) (models.Category, error)
	UpdateCategory(input FormUpdateCategory) (models.Category, error)
	DeleteCategory(input InputCategoryID) error
}
