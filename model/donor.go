package model

import "time"

// Patient defines comment model
type Donor struct {
	ID        string
	UserID    string
	Patient   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
