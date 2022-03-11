package model

import "time"

// Patient defines comment model
type BloodRequest struct {
	ID        string
	UserID    string
	Patient   string
	Request   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
