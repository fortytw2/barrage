package config

import (
  "github.com/BurntSushi/toml"
)

type config struct {
  SourceFolder  string `toml:"inputfolder"`
  StorageFolder string `toml:"storagefolder"`
  DB            string `toml:"database"`
  Port          string 
}

var Config config

func init() {
  if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
    panic(err)
  }
}
