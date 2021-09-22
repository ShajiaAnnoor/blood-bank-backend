package cfg

import "github.com/spf13/viper"

//Redis struct stores the redis related configuration information
type Redis struct {
	Addr     string
	Password string
	DB       int
}

//LoadRedis provides redis related configuration details
func LoadRedis() Redis {
	return Redis{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
}
