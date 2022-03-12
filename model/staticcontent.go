package model

import "time"

// Patient defines static content model
type StaticContent struct {
	ID            string
	UserID        string
	StaticContent string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
