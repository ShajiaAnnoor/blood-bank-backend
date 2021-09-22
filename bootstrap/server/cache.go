package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/cache/redis"
	"gitlab.com/Aubichol/blood-bank-backend/container"
)

func Cache(c container.Container) {
	c.Register(redis.NewSession)
	c.Register(redis.NewConnectionStatus)
}
