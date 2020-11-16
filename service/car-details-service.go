package service

// CarDetailsService represents the car details
type CarDetailsService interface {
	GetDetails()
}

type carService struct{}

// NewCarDetailsService creates a new instance of CarDetailsService
func NewCarDetailsService() CarDetailsService {
	return &carService{}
}

// GetDetails get details from car API
func (service *carService) GetDetails() {

}
