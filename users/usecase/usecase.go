package usecase

import (
	"errors"
	"payment/models"
	"payment/users"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	UserRepo users.Repository
}

func NewUserUsecase(repo users.Repository) *userUsecase {
	return &userUsecase{
		UserRepo: repo,
	}
}

func (u *userUsecase) RegisterUser(input users.RegisterUserInput) (models.User, error) {
	user := models.User{}
	user.Username = input.Username
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)

	newUser, err := u.UserRepo.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *userUsecase) Login(input users.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password
	user, err := u.UserRepo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found with that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUsecase) ChangePassword(input users.ChangePasswordInput) (models.User, error) {
	email := input.Email
	password := input.Password
	user, err := u.UserRepo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found with that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(newPassword)
	updatedUser, err := u.UserRepo.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (u *userUsecase) GetUserByID(ID int) (models.User, error) {
	user, err := u.UserRepo.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found found with that ID")
	}

	return user, nil
}
