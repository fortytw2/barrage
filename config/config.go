package config

import (
	"github.com/BurntSushi/toml"
)

type deserializedConfig struct {
	VideoFolder string `toml:"videofolder"`
	Port        string
}

var VideoFolder string
var Port string

func init() {

	var config deserializedConfig

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	VideoFolder = config.VideoFolder
	Port = config.Port

}
