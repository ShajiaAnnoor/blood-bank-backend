package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gitlab.com/Aubichol/blood-bank-backend/token"
)

//Token registers token related providers
func Token(c container.Container) {
	c.Register(token.NewRegisterStore)
}
