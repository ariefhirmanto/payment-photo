package repository

import (
	"payment/models"

	"gorm.io/gorm"
)

type frameRepository struct {
	frameDB *gorm.DB
}

func NewFrameRepository(db *gorm.DB) *frameRepository {
	return &frameRepository{frameDB: db}
}

func (r *frameRepository) CreateFrameImage(input models.Frame) (models.Frame, error) {
	err := r.frameDB.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

func (r *frameRepository) GetByID(ID int64) (models.Frame, error) {
	var frame models.Frame
	// with index
	// err := r.promoDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&promoCode).Error
	// without indexing
	err := r.frameDB.Where("id = ?", ID).Find(&frame).Error
	if err != nil {
		return frame, err
	}

	return frame, nil
}

func (r *frameRepository) GetAll() ([]models.Frame, error) {
	var frames []models.Frame

	err := r.frameDB.Find(&frames).Error
	if err != nil {
		return frames, err
	}

	return frames, nil
}

func (r *frameRepository) GetByCategoryID(input int64) ([]models.Frame, error) {
	var frames []models.Frame
	// with index
	// err := r.promoDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&promoCode).Error
	// without indexing
	err := r.frameDB.Where("category_id = ?", input).Where("available = 1").Find(&frames).Error
	if err != nil {
		return frames, err
	}

	return frames, nil
}

func (r *frameRepository) Delete(input models.Frame) (bool, error) {
	err := r.frameDB.Delete(&input).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *frameRepository) Update(input models.Frame) (models.Frame, error) {
	err := r.frameDB.Omit("created_at").Save(&input).Error
	if err != nil {
		return input, err
	}
	return input, nil
}

func (r *frameRepository) GetFrameByLocation(input string) ([]models.Frame, error) {
	var frames []models.Frame
	// with index
	// err := r.promoDB.Clauses(hints.UseIndex("idx_status")).Where("id = ?", ID).Find(&promoCode).Error
	// without indexing
	err := r.frameDB.Where("location = ?", input).Where("available = 1").Find(&frames).Error
	if err != nil {
		return frames, err
	}

	return frames, nil
}
