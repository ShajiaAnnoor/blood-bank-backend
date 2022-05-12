package model

import "time"

// Notice defines notice model
type Notice struct {
	ID          string
	PatientName string
	BloodGroup  string
	District    string
	Address     string
	UserID      string
	Description string
	Title       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
