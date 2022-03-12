package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/cfg"
	"gitlab.com/Aubichol/blood-bank-backend/container"
)

//Cfg registers configuration related providers
func Cfg(c container.Container) {
	//LoadSession provides a constructor to dig container that loads the duration of a session and the maximum number of sessions a user can have
	c.Register(cfg.LoadSession)
	//LoadMongo provides a container that loads mongodb server url name and the name of the database
	c.Register(cfg.LoadMongo)
	//LoadRedis provides a contructor to dig container for loading redis related configurations
	c.Register(cfg.LoadRedis)
	//LoadServer provides a consturctor to dig container for loading go server related configurations
	c.Register(cfg.LoadServer)
	//LoadConnectionCache provides a constructor to dig container for connection cache
	//TODO: What is connection cache?
	c.Register(cfg.LoadConnectionCache)
	//LoadLimit provides a constructor to dig container for limit related data
	c.Register(cfg.LoadLimit)
	//LoadWSClient provides a construction to dig container for Web Socket client
	c.Register(cfg.LoadWSClient)
	//LoadFile provides a constructor to dig container for maintaining file
	c.Register(cfg.LoadFile)
}
