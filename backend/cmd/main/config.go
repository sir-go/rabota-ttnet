package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"main/internal/server"
	"main/internal/store"
)

type Config struct {
	Service server.Config `toml:"service"`
	Db      store.Config  `toml:"db"`
}

func unmarshal(b []byte) (cfg *Config, err error) {
	if err != nil {
		return
	}
	cfg = new(Config)
	if b == nil {
		return nil, errors.New("config data is empty")
	}
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
