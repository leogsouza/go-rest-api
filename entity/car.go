package entity

// Car represents a car entity
type Car struct {
	CarData `json:"car"`
}

// CarData represents car details data
type CarData struct {
	ID    int    `json:"id"`
	Brand string `json:"car"`
	Model string `json:"car_model"`
	Year  int    `json:"car_model_year"`
}
