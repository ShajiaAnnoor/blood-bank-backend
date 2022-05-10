package model

import "time"

// Donor defines donor model
type Donor struct {
	ID           string
	UserID       string
	Patient      string
	Phone        string
	District     string
	BloodGroup   string
	Address      string
	Availability bool
	TimesDonated int
	Name         string
	IsDeleted    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
