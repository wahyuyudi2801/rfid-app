package models

import "time"

type RfidCard struct {
	ID             uint      `json:"id"`
	RfidTag        string    `json:"rfid_tag"`
	ActivationDate time.Time `json:"activation_date"`
	Status         int8      `json:"status"`
	EmployeeID     uint      `json:"employee_id"`
}
