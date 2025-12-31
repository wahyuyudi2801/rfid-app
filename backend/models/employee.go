package models

type Employee struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	RegistrationNumber string `json:"registration_number"`
	Position           string `json:"position"`
	WorkUnit           string `json:"work_unit"`
	Status             int    `json:"status"`
}
