package service

import (
	"fmt"
	"net/http"
)

// OwnerService represents the car fetch data service
type OwnerService interface {
	FetchData()
}

const (
	ownerServiceURL = "https://myfakeapi.com/api/users/1"
)

type ownerService struct{}

// NewOwnerService creates a new CarService instance
func NewOwnerService() OwnerService {
	return &carService{}
}

// FetchData fetches the owner data from specific endpoint
func (c *ownerService) FetchData() {
	client := http.Client{}

	fmt.Printf("Fetching the url %s", ownerServiceURL)

	// Call the external API
	resp, _ := client.Get(carServiceURL)
	fmt.Println("Response", resp)

	// TODO: Write response to the channel

}
