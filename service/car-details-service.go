package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/leogsouza/go-rest-api/entity"
)

var (
	carServ   CarService   = NewCarService()
	ownerServ OwnerService = NewOwnerService()

	carDataChannel   = make(chan *http.Response)
	ownerDataChannel = make(chan *http.Response)
)

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
	go carServ.FetchData()
	// go routine call endpoint 2 get data from https://myfakeapi.com/api/users/1
	go ownerServ.FetchData()

	car, _ := getCarData()
	owner, _ := getOwnerdata()

	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}

	// create carChannel to get the data from endpoint 1
	// create ownerChannel to get the data from endpoint 2

}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel

	var car entity.Car

	err := json.NewDecoder(r1.Body).Decode(&car)

	if err != nil {
		fmt.Println(err.Error())
		return car, err
	}

	return car, nil
}

func getOwnerdata() (entity.Owner, error) {
	r1 := <-ownerDataChannel

	var owner entity.Owner

	err := json.NewDecoder(r1.Body).Decode(&owner)

	if err != nil {
		fmt.Println(err.Error())
		return owner, err
	}

	return owner, nil
}
