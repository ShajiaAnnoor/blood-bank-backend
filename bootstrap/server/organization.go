package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gitlab.com/Aubichol/blood-bank-backend/organization"
)

//Organization registers organization related providers
func Organization(c container.Container) {
	c.Register(organization.NewCreate)
	c.Register(organization.NewReader)
	c.Register(organization.NewUpdate)
	c.Register(organization.NewDelete)
}
