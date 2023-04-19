package usecase

import (
	"fmt"
	"log"
	"os"
	frame "payment/frame"
	category "payment/frame_category"
	"payment/models"
	"strings"
)

type frameUsecase struct {
	FrameRepo    frame.Repository
	CategoryRepo category.Repository
	BaseUrl      string
}

func NewFrameUsecase(repo frame.Repository, categoryRepo category.Repository, baseUrl string) *frameUsecase {
	return &frameUsecase{
		FrameRepo:    repo,
		CategoryRepo: categoryRepo,
		BaseUrl:      baseUrl,
	}
}

func (u *frameUsecase) SaveFrameImage(input frame.FormInputFrame, fileLocation string) (models.Frame, error) {
	category, err := u.CategoryRepo.GetByID(input.CategoryID)
	if err != nil {
		return models.Frame{}, err
	}

	Url := strings.Replace(u.BaseUrl+fileLocation, "//app", "", 1)
	saveFrame := models.Frame{}
	saveFrame.CategoryID = input.CategoryID
	saveFrame.Category = category
	saveFrame.Location = input.Location
	saveFrame.Name = input.Name
	saveFrame.Url = Url
	saveFrame.Counter = input.Counter
	saveFrame.Available = true
	log.Printf("[Frame][Usecase][SaveFrameImage] saveFrame request %+v", saveFrame)

	newFrame, err := u.FrameRepo.CreateFrameImage(saveFrame)
	if err != nil {
		return newFrame, err
	}

	return newFrame, nil
}

func (u *frameUsecase) GetFrameByID(input frame.InputFrameID) (models.Frame, error) {
	frame, err := u.FrameRepo.GetByID(input.ID)
	if err != nil {
		return models.Frame{}, err
	}

	return frame, nil
}

func (u *frameUsecase) GetFrameByName(input frame.InputFrameName) (models.Frame, error) {
	frame, err := u.FrameRepo.GetByName(input.Name)
	if (err != nil || frame == models.Frame{}) {
		return models.Frame{}, err
	}

	return frame, nil
}

func (u *frameUsecase) GetFrameByCategoryID(input frame.InputCategoryID) ([]models.Frame, error) {
	frames, err := u.FrameRepo.GetByCategoryID(input.ID)
	if err != nil {
		return frames, err
	}

	return frames, nil
}

func (u *frameUsecase) DeleteFrame(input frame.InputFrameID) error {
	frame, err := u.FrameRepo.GetByID(input.ID)
	if err != nil {
		return err
	}

	directory := strings.Replace(frame.Url, u.BaseUrl, "/app/", 1)
	err = os.Remove(directory)
	if err != nil {
		log.Printf("[Frame][Usecase][DeleteFrame] Error deleting frame %+v", err)
		return err
	}

	_, err = u.FrameRepo.Delete(frame)
	if err != nil {
		return err
	}

	return nil
}

func (u *frameUsecase) ChangeStatusFrame(input frame.InputFrameID) error {
	frame, err := u.FrameRepo.GetByID(input.ID)
	if err != nil {
		return err
	}

	statusAvailable := frame.Available
	frame.Available = !statusAvailable
	_, err = u.FrameRepo.Update(frame)
	if err != nil {
		return err
	}

	return nil
}

func (u *frameUsecase) GetAllFrame() ([]models.Frame, error) {
	frames, err := u.FrameRepo.GetAll()
	fmt.Printf("Frame: %+v\n", frames)
	if err != nil {
		log.Printf("[Frame][Usecase][GetAllFrame] Error get frame: %+v", err)
		return frames, err
	}

	return frames, nil
}

func (u *frameUsecase) GetFrameByCategoryName(input frame.InputCategoryName) ([]models.Frame, error) {
	category, err := u.CategoryRepo.GetByName(input.CategoryName)
	if err != nil {
		return []models.Frame{}, err
	}

	frames, err := u.FrameRepo.GetByCategoryID(category.ID)
	if err != nil {
		return frames, err
	}

	return frames, nil
}

func (u *frameUsecase) GetFrameByLocation(input frame.InputLocationName) ([]models.Frame, error) {
	frames, err := u.FrameRepo.GetFrameByLocation(input.Location)
	if err != nil {
		return frames, err
	}

	return frames, nil
}
