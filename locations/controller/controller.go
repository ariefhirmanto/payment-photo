package controller

import (
	location "payment/locations"
)

type locationController struct {
	locationUC location.Usecase
}

func NewLocationController(locationUC location.Usecase) *locationController {
	return &locationController{locationUC: locationUC}
}
