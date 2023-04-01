package repository

import (
	"payment/models"

	"gorm.io/gorm"
)

type categoryRepository struct {
	categoryDB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{categoryDB: db}
}

func (r *categoryRepository) CreateCategory(input models.Category) (models.Category, error) {
	err := r.categoryDB.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

func (r *categoryRepository) Update(input models.Category) (models.Category, error) {
	err := r.categoryDB.Omit("created_at").Save(&input).Error
	if err != nil {
		return input, err
	}
	return input, nil
}

func (r *categoryRepository) GetByID(ID int64) (models.Category, error) {
	var category models.Category
	// with index
	// err := r.promoDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&promoCode).Error
	// without indexing
	err := r.categoryDB.Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	err := r.categoryDB.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *categoryRepository) GetByName(input string) (models.Category, error) {
	var category models.Category
	// with index
	// err := r.promoDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&promoCode).Error
	// without indexing
	err := r.categoryDB.Where("name = ?", input).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) Delete(input models.Category) (bool, error) {
	err := r.categoryDB.Delete(&input).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
