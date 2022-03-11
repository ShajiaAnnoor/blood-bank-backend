package model

import "time"

// Patient defines comment model
type StaticContent struct {
	ID            string
	UserID        string
	StaticContent string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
