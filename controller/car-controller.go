package controller

import (
	"encoding/json"
	"net/http"

	"github.com/leogsouza/go-rest-api/service"
)

// CarController holds the methods to handle car requests
type CarController interface {
	GetCarDetails(w http.ResponseWriter, r *http.Request)
}

type carController struct {
	service service.CarDetailsService
}

// NewCarController creates a new CarController instance
func NewCarController(serv service.CarDetailsService) CarController {
	return &carController{
		service: serv,
	}
}

func (c *carController) GetCarDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := c.service.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
