package model

import "time"

// Patient defines comment model
type Patient struct {
	ID        string
	UserID    string
	StatusID  string
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
