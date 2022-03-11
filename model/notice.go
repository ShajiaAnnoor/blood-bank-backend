package model

import "time"

// Patient defines comment model
type Notice struct {
	ID        string
	UserID    string
	Patient   string
	Notice    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
