package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent"
)

//StaticContent registers staticcontent related providers
func StaticContent(c container.Container) {
	c.Register(staticcontent.NewCreate)
	c.Register(staticcontent.NewReader)
	c.Register(staticcontent.NewUpdate)
	c.Register(staticcontent.NewDelete)
}
