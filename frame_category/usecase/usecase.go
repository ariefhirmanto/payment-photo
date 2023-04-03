package usecase

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	category "payment/frame_category"
	"payment/models"
)

type categoryUsecase struct {
	CategoryRepo category.Repository
	Env          string
}

func NewCategoryUsecase(repo category.Repository, env string) *categoryUsecase {
	return &categoryUsecase{
		CategoryRepo: repo,
		Env:          env,
	}
}

func (u *categoryUsecase) CreateCategory(input category.FormInputCategory) (models.Category, error) {
	category := models.Category{}
	category.Name = input.Name
	category.InterRowPadding = input.InterRowPadding
	category.TopFramePadding = input.TopFramePadding
	category.InterColPadding = input.InterColPadding
	category.CustomPadding = input.CustomPadding
	category.ImageID = input.ImageID
	category.Width = input.Width
	category.Height = input.Height

	categoryFromDB, _ := u.CategoryRepo.GetByName(input.Name)
	if (models.Category{}) != categoryFromDB {
		log.Printf("[Category][Usecase][CreateCategory] Error category already exists %+v", input.Name)
		return category, errors.New("category already exists")
	}

	parentDirectory := filepath.Join(getDirectory(u.Env), "images")
	err := os.Mkdir(filepath.Join(parentDirectory, input.Name), os.ModePerm)
	if err != nil {
		log.Printf("[Category][Usecase][CreateCategory] Error creating category %+v", err)
		return category, err
	}

	savedCategory, err := u.CategoryRepo.CreateCategory(category)
	if err != nil {
		log.Printf("[Category][Usecase][CreateCategory] Error creating category %+v", err)
		return savedCategory, err
	}

	log.Printf("[Category][Usecase][CreateCategory] Success create category %+v", savedCategory)
	return savedCategory, nil
}

func (u *categoryUsecase) GetAllCategory() ([]models.Category, error) {
	categories, err := u.CategoryRepo.GetAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (u *categoryUsecase) FindByIDForFrame(input int64) (models.Category, error) {
	category, err := u.CategoryRepo.GetByID(input)
	if err != nil {
		log.Printf("[Promo][Usecase][FindByID] Error get promo transaction with ID %+v", input)
		return category, err
	}

	return category, nil
}

func (u *categoryUsecase) FindByID(input category.InputCategoryID) (models.Category, error) {
	category, err := u.CategoryRepo.GetByID(input.ID)
	if err != nil {
		log.Printf("[Promo][Usecase][FindByID] Error get promo transaction with ID %+v", input.ID)
		return category, err
	}

	return category, nil
}

func (u *categoryUsecase) FindByName(input category.InputCategoryName) (models.Category, error) {
	category, err := u.CategoryRepo.GetByName(input.Name)
	if err != nil {
		log.Printf("[Promo][Usecase][FindByPromoCode] Error get promo transaction with code %+v", input.Name)
		return category, err
	}

	return category, nil
}

func (u *categoryUsecase) UpdateCategory(input category.FormUpdateCategory) (models.Category, error) {
	category, err := u.CategoryRepo.GetByID(input.ID)
	if err != nil {
		log.Printf("[Promo][Usecase][ClaimPromo] Error get promo transaction with code %+v", input.ID)
		return category, err
	}

	updateData := models.Category{}
	updateData.Name = input.Name
	updateData.ID = input.ID
	if input.InterRowPadding != 0 {
		updateData.InterRowPadding = input.InterRowPadding
	}

	if input.InterColPadding != 0 {
		updateData.InterColPadding = input.InterColPadding
	}

	if input.TopFramePadding != 0 {
		updateData.TopFramePadding = input.TopFramePadding
	}

	if input.CustomPadding != 0 {
		updateData.CustomPadding = input.CustomPadding
	}

	if input.ImageID != 0 {
		updateData.ImageID = input.ImageID
	}

	if input.Width != 0 {
		updateData.Width = input.Width
	}

	if input.Height != 0 {
		updateData.Height = input.Height
	}

	err = os.Rename(filepath.Join("images", category.Name), filepath.Join("images", input.Name))
	if err != nil {
		log.Printf("[Promo][Usecase][ClaimPromo] Failed to update promo with code %+v", input.Name)
		return category, err
	}

	newCategory, err := u.CategoryRepo.Update(updateData)
	if err != nil {
		log.Printf("[Promo][Usecase][ClaimPromo] Failed to update promo with code %+v", input.Name)
		return newCategory, err
	}

	log.Printf("[Promo][Usecase][ClaimPromo] Success claim promo %+v", newCategory)
	return newCategory, nil
}

func (u *categoryUsecase) DeleteCategory(input category.InputCategoryID) error {
	category, err := u.CategoryRepo.GetByID(input.ID)
	if err != nil {
		log.Printf("[Promo][Usecase][DeletePromo] Error get promo transaction with ID %+v", input.ID)
		return err
	}

	parentDirectory := filepath.Join(getDirectory(u.Env), "images")
	err = os.RemoveAll(filepath.Join(parentDirectory, category.Name))
	if err != nil {
		log.Printf("[Category][Usecase][CreateCategory] Error creating category %+v", err)
		return err
	}

	isSuccess, err := u.CategoryRepo.Delete(category)
	if err != nil {
		log.Printf("[Promo][Usecase][DeletePromo] Failed to update promo with ID %+v", input.ID)
		return err
	}

	log.Printf("[Promo][Usecase][DeletePromo] Success delete promo %+v", isSuccess)
	return nil
}

func getDirectory(env string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	url := filepath.Dir(wd)
	if env != "local" {
		url = "/app" + filepath.Dir(wd)
	}

	return url
}
