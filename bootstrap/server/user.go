package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gitlab.com/Aubichol/blood-bank-backend/user"
)

//User registers user related providers
func User(c container.Container) {
	c.Register(user.NewRegistry)
	c.Register(user.NewSearcher)
	c.Register(user.NewLogin)
	c.Register(user.NewEmailAndPasswordChecker)
	c.Register(user.NewSessionVerifier)
	c.Register(user.NewMyProfile)
	//	c.Register(user.NewProfileUpdater)
	//	c.Register(user.NewProfilePicSaver)
}
