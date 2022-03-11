package model

import "time"

// Token defines token model
type Token struct {
	ID        string
	UserID    string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
