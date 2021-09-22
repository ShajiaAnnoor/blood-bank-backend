package cfg

import "github.com/spf13/viper"

//Mongo stores mongodb related configuration information
type Mongo struct {
	URI      string
	Database string
}

//LoadMongo provides mongodb related configuration to server
func LoadMongo() Mongo {
	return Mongo{
		URI:      viper.GetString("mongo.uri"),
		Database: viper.GetString("mongo.database"),
	}
}
