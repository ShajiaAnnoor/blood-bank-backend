package cfg

import "github.com/spf13/viper"

//File struct stores file related constructor information
type File struct {
	// Loc is the location where the files are stored
	Loc string
	// MaxSize is the maximum file size that can be stored.
	MaxSize int64
}

//LoadFile provides file related information
func LoadFile() File {
	return File{
		Loc:     viper.GetString("file.location"),
		MaxSize: viper.GetInt64("file.max_size"),
	}
}
