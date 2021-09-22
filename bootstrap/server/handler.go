package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/api"
	"gitlab.com/Aubichol/blood-bank-backend/container"
)

//Handler registers provider that returns root handler
func Handler(c container.Container) {
	c.Register(api.Handler)
}
