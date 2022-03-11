package model

import "time"

// Patient defines comment model
type Organization struct {
	ID           string
	UserID       string
	Patient      string
	Organization string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
