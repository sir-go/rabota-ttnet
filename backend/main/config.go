package main

import (
	"flag"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type (
	Duration struct {
		time.Duration
	}

	CfgService struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Secret   string `toml:"secret"`
		Static   string `toml:"static"`
		Timeouts struct {
			Write *Duration `toml:"write"`
			Read  *Duration `toml:"read"`
			Idle  *Duration `toml:"idle"`
		} `toml:"timeouts"`
		RunMode string `toml:"run_mode"`
	}

	CfgDb struct {
		Host     string    `toml:"host"`
		Port     int       `toml:"port"`
		User     string    `toml:"user"`
		Password string    `toml:"password"`
		DbName   string    `toml:"dbname"`
		Timeout  *Duration `toml:"timeout"`
		IdLen    int       `toml:"id_len"`
		IdTries  int       `toml:"id_tries"`
	}

	Config struct {
		Service CfgService `toml:"service"`
		Db      CfgDb      `toml:"db"`
		Path    string
	}
)

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func ConfigInit() *Config {
	fCfgPath := flag.String("c", DefaultConfFile, "path to conf file")
	flag.Parse()

	conf := new(Config)
	file, err := os.Open(*fCfgPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if file == nil {
			return
		}
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err = toml.DecodeFile(*fCfgPath, &conf); err != nil {
		panic(err)
	}
	conf.Path = *fCfgPath
	return conf
}
