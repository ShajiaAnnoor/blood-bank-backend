package model

import "time"

// BloodRequest defines blood request model
type BloodRequest struct {
	ID         string
	UserID     string
	Patient    string
	Request    string
	BloodGroup string
	IsDeleted  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
