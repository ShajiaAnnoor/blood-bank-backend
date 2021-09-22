package cfg

import "github.com/spf13/viper"

//Server is a struct that stores the go server related configuration information
type Server struct {
	Port string
}

//LoadServer returns a server instance to it's calling library.
func LoadServer() Server {
	return Server{
		//Port number of the server is retrieved from the configuration file using built-in viper library
		Port: viper.GetString("server.port"),
	}
}
