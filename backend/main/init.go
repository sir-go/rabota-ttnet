package main

import (
	"log"
	"os"
	"os/signal"
	//"regexp"
)

var (
	CFG *Config
	LOG *log.Logger
	DBC *DB
	SRV *Server
)

func initInterrupt() {
	LOG.Println("-- start --")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(c chan os.Signal) {
		for range c {
			SRV.Shutdown()
			LOG.Println("-- stop --")
			os.Exit(137)
		}
	}(c)
}

func init() {
	var err error
	CFG = ConfigInit()
	LOG = initLogging()
	initInterrupt()
	LOG.Printf("config path: %s", CFG.Path)
	DBC = new(DB)
	SRV = new(Server)
	//RxpNumbersOnly, err = regexp.Compile(`[^0-9]+`)
	eh(err)
}
