package config

import (
	"github.com/BurntSushi/toml"
)

type deserializedConfig struct {
	SourceFolder  string `toml:"sourcefolder"`
	StorageFolder string `toml:"storagefolder"`
	DB            string `toml:"database"`
	Port          string
}

var SourceFolder string
var StorageFolder string
var DB string
var Port string

func init() {

	var config deserializedConfig

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	SourceFolder = config.SourceFolder
	StorageFolder = config.StorageFolder
	DB = config.DB
	Port = config.Port

}
