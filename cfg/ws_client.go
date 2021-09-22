package cfg

import (
	"time"

	"github.com/spf13/viper"
)

//WSClient struct holds and provides web socket related configuration information
type WSClient struct {
	ReadQBuffer     int
	WriteQBuffer    int
	WriteWait       time.Duration
	PongWait        time.Duration
	PingPeriod      time.Duration
	ReadMessageSize int64
}

//LoadWSClient provides web socket related configuration information
func LoadWSClient() WSClient {

	return WSClient{
		ReadQBuffer:     viper.GetInt("wsclient.readq_buffer"),
		WriteQBuffer:    viper.GetInt("wsclient.writeq_buffer"),
		PongWait:        viper.GetDuration("wsclient.pong_wait"),
		PingPeriod:      viper.GetDuration("wsclient.ping_period"),
		WriteWait:       viper.GetDuration("wsclient.write_wait"),
		ReadMessageSize: viper.GetInt64("wsclient.read_message_size"),
	}
}
