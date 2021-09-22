package server

import "github.com/spf13/viper"

//Viper configures viper related configuration details
func Viper() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./resource/config.yaml")
	return viper.ReadInConfig()
}
