package server

import (
	"gitlab.com/Aubichol/hrishi-backend/container"
	"gitlab.com/Aubichol/hrishi-backend/ws"
)

//WS regisers web socket related providers
func WS(c container.Container) {
	c.Register(ws.NewHub)
	c.RegisterWithName(ws.NewAuthHandler, "auth")
	c.Register(ws.NewClientStore)
}
