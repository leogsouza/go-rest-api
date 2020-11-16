package service

import (
	"fmt"
	"net/http"
)

// CarService represents the car fetch data service
type CarService interface {
	FetchData()
}

const (
	carServiceURL = "https://myfakeapi.com/api/cars/1"
)

type carService struct{}

// NewCarService creates a new CarService instance
func NewCarService() CarService {
	return &carService{}
}

// FetchData fetches the car data from specific endpoint
func (c *carService) FetchData() {
	client := http.Client{}

	fmt.Printf("Fetching the url %s", carServiceURL)

	// Call the external API
	resp, _ := client.Get(carServiceURL)
	fmt.Println("Response", resp)
	// TODO: Write response to the channel

}
