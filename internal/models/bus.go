package models

import "time"

type Bus struct {
	ID           int       `json:"id"`
	BusNumber    string    `json:"bus_number"`
	OperatorName string    `json:"operator_name"`
	TotalSeats   int       `json:"total_seats"`
	BusType      string    `json:"bus_type"`
	Amenities    []string  `json:"amenities"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
