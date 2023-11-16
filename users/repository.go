package users

import "payment/models"

type Repository interface {
	Save(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindByID(ID int) (models.User, error)
}
