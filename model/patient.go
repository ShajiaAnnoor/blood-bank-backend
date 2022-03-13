package model

import "time"

// Patient defines patient model
type Patient struct {
	ID        string
	UserID    string
	Patient   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
