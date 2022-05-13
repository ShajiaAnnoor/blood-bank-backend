package model

import "time"

// Organization defines organization model
type Organization struct {
	ID           string
	UserID       string
	Organization string
	Name         string
	Phone        string
	District     string
	Description  string
	Address      string
	IsDeleted    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
