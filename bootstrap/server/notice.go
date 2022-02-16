package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gitlab.com/Aubichol/blood-bank-backend/notice"
)

//Notice registers notice related providers
func Notice(c container.Container) {
	c.Register(notice.NewCreate)
	c.Register(notice.NewReader)
	c.Register(notice.NewUpdate)
	c.Register(notice.NewDelete)
}
