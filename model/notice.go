package model

import "time"

// Notice defines notice model
type Notice struct {
	ID        string
	UserID    string
	Patient   string
	Notice    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
