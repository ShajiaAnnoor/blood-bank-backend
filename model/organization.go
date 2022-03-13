package model

import "time"

// Organization defines organization model
type Organization struct {
	ID           string
	UserID       string
	Patient      string
	Organization string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
