package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"gitlab.com/Aubichol/blood-bank-backend/cache"
	"gitlab.com/Aubichol/blood-bank-backend/cfg"
	"gitlab.com/Aubichol/blood-bank-backend/pkg/tag"
)

const (
	connectionStatusPrefix string = "connection_status"
)

type connectionStatus struct {
	c   *redis.Client
	cfg cfg.ConnectionCache
}

func (c *connectionStatus) key(id1, id2 string) string {
	return fmt.Sprintf("%s_%s", connectionStatusPrefix, tag.Unique(id1, id2))
}

func (c *connectionStatus) Has(id1, id2 string) (bool, error) {
	result := c.c.Get(c.key(id1, id2))
	if err := result.Err(); err != nil {
		return false, err
	}

	v, err := result.Int()
	if err != nil {
		return false, err
	}

	return v == 1, nil
}

func (c *connectionStatus) Set(id1, id2 string, status bool) error {
	v := 1
	if !status {
		v = 0
	}

	result := c.c.Set(c.key(id1, id2), v, c.cfg.Length)
	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func NewConnectionStatus(c *redis.Client, cfg cfg.ConnectionCache) cache.ConnectionStatus {
	return &connectionStatus{c, cfg}
}
