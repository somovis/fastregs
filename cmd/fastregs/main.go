package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/somovis/fastregs/internal/app/fastregs"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/fastregs.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := fastregs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := fastregs.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
