package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"main/internal/server"
	"main/internal/store"
)

type Config struct {
	Service server.Config `toml:"service"`
	Db      store.Config  `toml:"db"`
	Path    string
}

func unmarshal(b []byte) (cfg *Config, err error) {
	if err != nil {
		return
	}
	cfg = new(Config)
	err = toml.Unmarshal(b, cfg)
	return
}

func LoadConfig(configPath string) (cfg *Config, err error) {
	b, err := ioutil.ReadFile(filepath.Clean(configPath))
	if err != nil {
		return nil, err
	}
	return unmarshal(b)
}
