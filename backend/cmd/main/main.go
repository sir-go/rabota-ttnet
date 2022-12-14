package main

import (
	"flag"

	zlog "github.com/rs/zerolog/log"

	"main/internal/server"
)

func eh(err error, msg string) {
	if err != nil {
		zlog.Err(err).Msg(msg)
		panic(err)
	}
}

func main() {
	fConfigPath := flag.String("c", "/opt/config.toml", "path to config file")
	cfg, err := LoadConfig(*fConfigPath)
	eh(err, "config reading")

	srv, err := server.Run(cfg.Service, cfg.Db)
	eh(err, "server starting")

	initInterrupt(srv.Close)
}
