package model

import "time"

// Donor defines donor model
type Donor struct {
	ID        string
	UserID    string
	Patient   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
