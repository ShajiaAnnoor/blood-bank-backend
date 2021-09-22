package server

import (
	"gitlab.com/Aubichol/hrishi-backend/api"
	"gitlab.com/Aubichol/hrishi-backend/container"
)

//Handler registers provider that returns root handler
func Handler(c container.Container) {
	c.Register(api.Handler)
}
