package server

import (
	"time"
)

type Config struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Secret   string `toml:"secret"`
	Static   string `toml:"static"`
	Timeouts struct {
		Write time.Duration `toml:"write"`
		Read  time.Duration `toml:"read"`
		Idle  time.Duration `toml:"idle"`
	} `toml:"timeouts"`
	RunMode string `toml:"run_mode"`
}
