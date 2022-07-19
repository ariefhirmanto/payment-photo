package repository

import (
	"payment/models"

	"gorm.io/gorm"
)

type userRepository struct {
	UserDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{UserDB: db}
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	err := r.UserDB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.UserDB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID int) (models.User, error) {
	var user models.User

	err := r.UserDB.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User

	err := r.UserDB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
