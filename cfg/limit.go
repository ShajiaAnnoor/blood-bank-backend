package cfg

import (
	"github.com/spf13/viper"
)

// Limit is a struct that stores the limits regarding some of the application level features
type Limit struct {
	FriendList    int
	ChatList      int
	MessageLength int
}

//LoadLimit returns an instance for defining some application specific limits to it's calling library
func LoadLimit() Limit {
	return Limit{
		FriendList:    viper.GetInt("limit.friend_list"),
		ChatList:      viper.GetInt("limit.chat_list"),
		MessageLength: viper.GetInt("limit.message_length"),
	}
}
