package cfg

import (
	"time"

	"github.com/spf13/viper"
)

// Session holds config for session
type Session struct {
	// defines length of a session
	Length time.Duration
	// defines maximum number of session per user
	MaxPerUser int
}

//LoadSession providies session related information
func LoadSession() Session {
	viper.GetDuration("session.length")

	return Session{
		Length:     viper.GetDuration("session.length"),
		MaxPerUser: viper.GetInt("session.max_per_user"),
	}
}
