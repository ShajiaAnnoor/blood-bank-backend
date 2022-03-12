package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gitlab.com/Aubichol/blood-bank-backend/patient"
)

//Patient registers patient related providers
func Patient(c container.Container) {
	c.Register(patient.NewCreate)
	c.Register(patient.NewReader)
	c.Register(patient.NewUpdate)
	c.Register(patient.NewDelete)
}
