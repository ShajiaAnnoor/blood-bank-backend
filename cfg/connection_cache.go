package cfg

import (
	"time"

	"github.com/spf13/viper"
)

//ConnectionCache stores the configurations loaded from the configuration files about the connection cache
type ConnectionCache struct {
	Length time.Duration
}

//LoadConnectionCache returns a connection cache instance to it's calling library.
func LoadConnectionCache() ConnectionCache {
	return ConnectionCache{
		Length: viper.GetDuration("connection_cache.length"),
	}
}
