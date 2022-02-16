package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gitlab.com/Aubichol/blood-bank-backend/donor"
)

//Donor registers donor related providers
func Donor(c container.Container) {
	c.Register(donor.NewCreate)
	c.Register(donor.NewReader)
	c.Register(donor.NewUpdate)
	c.Register(donor.NewDelete)
}
