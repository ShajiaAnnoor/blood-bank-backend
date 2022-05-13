package model

import "time"

// Patient defines patient model
type Patient struct {
	ID         string
	UserID     string
	Name       string
	BloodGroup string
	District   string
	Phone      string
	Address    string
	IsDeleted  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
