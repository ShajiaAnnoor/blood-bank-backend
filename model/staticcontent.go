package model

import "time"

// StaticContent defines static content model
type StaticContent struct {
	ID        string
	UserID    string
	Text      string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
