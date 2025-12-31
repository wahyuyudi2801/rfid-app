package models

import "time"

type Attendance struct {
	ID               uint      `json:"id"`
	TapType          int8      `json:"tap_type"`
	ValidationStatus int8      `json:"validation_status"`
	CreatedAt        time.Time `json:"created_at"`
	RfidID           uint      `json:"rfid_id"`
	LocationID       uint      `json:"location_id"`
}
