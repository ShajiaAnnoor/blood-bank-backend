package server

import (
	"gitlab.com/Aubichol/hrishi-backend/cache/redis"
	"gitlab.com/Aubichol/hrishi-backend/container"
)

func Cache(c container.Container) {
	c.Register(redis.NewSession)
	c.Register(redis.NewConnectionStatus)
}
