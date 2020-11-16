package service

import "github.com/leogsouza/go-rest-api/entity"

// CarDetailsService represents the car details
type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type carDetailsService struct{}

// NewCarDetailsService creates a new instance of CarDetailsService
func NewCarDetailsService() CarDetailsService {
	return &carDetailsService{}
}

// GetDetails get details from car API
func (service *carDetailsService) GetDetails() entity.CarDetails {
	// go routine call endpoint 1
	// go routine call endpoint 2 get data from https://myfakeapi.com/api/users/1

	// create carChannel to get the data from endpoint 1
	// create ownerChannel to get the data from endpoint 2
	var carDetails entity.CarDetails

	return carDetails
}
