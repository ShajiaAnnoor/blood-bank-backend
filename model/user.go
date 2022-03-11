package model

import "time"

// User defines user model
type User struct {
	ID         string
	FirstName  string
	LastName   string
	Gender     string
	BirthDate  BirthDate
	Email      string
	Password   string
	Verified   bool
	Profile    map[string]interface{}
	ProfilePic string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//BirthDate defines birthdate model
type BirthDate struct {
	Year  int
	Month int
	Day   int
}
