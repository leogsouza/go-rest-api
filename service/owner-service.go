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
	return &ownerService{}
}

// FetchData fetches the owner data from specific endpoint
func (c *ownerService) FetchData() {
	client := http.Client{}

	fmt.Printf("Fetching the url %s", ownerServiceURL)

	// Call the external API
	resp, _ := client.Get(ownerServiceURL)
	fmt.Println("Response", resp)

	// Write response to the channel
	ownerDataChannel <- resp

}
