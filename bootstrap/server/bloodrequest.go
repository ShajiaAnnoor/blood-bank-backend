package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest"
	"gitlab.com/Aubichol/blood-bank-backend/container"
)

//Blood Request registers blood request related providers
func BloodRequest(c container.Container) {
	c.Register(bloodrequest.NewCreate)
	c.Register(bloodrequest.NewReader)
	c.Register(bloodrequest.NewUpdate)
	c.Register(bloodrequest.NewDelete)
}
