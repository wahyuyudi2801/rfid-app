package models

type Location struct {
	ID           uint   `json:"id"`
	LocationName string `json:"location_name"`
	ReaderCode   string `json:"reader_code"`
}
