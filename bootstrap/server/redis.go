package server

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/cfg"
	"gitlab.com/Aubichol/hrishi-backend/container"
)

//Redis provides a constructor for creating redis.Client instance from cfg.Redis config details
func Redis(c container.Container) {
	c.Register(func(cfg cfg.Redis) *redis.Client {
		client := redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB, // use default DB
		})

		if _, err := client.Ping().Result(); err != nil {
			logrus.Fatal(err)
		}

		return client
	})
}
