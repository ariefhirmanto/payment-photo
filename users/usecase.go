package users

import "payment/models"

type Usecase interface {
	RegisterUser(input RegisterUserInput) (models.User, error)
	Login(input LoginInput) (models.User, error)
	ChangePassword(input ChangePasswordInput) (models.User, error)
	GetUserByID(ID int) (models.User, error)
}
