package store

import (
	"time"
)

type Config struct {
	Host     string        `toml:"host"`
	Port     int           `toml:"port"`
	User     string        `toml:"user"`
	Password string        `toml:"password"`
	DbName   string        `toml:"dbname"`
	Timeout  time.Duration `toml:"timeout"`
	IdLen    int           `toml:"id_len"`
	IdTries  int           `toml:"id_tries"`
}
