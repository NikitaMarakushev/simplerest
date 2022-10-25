package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"simplerest/internal/app/apiserver"
)

var (
	configPath    string
	pathToLogFile = "configs/apiserver.toml"
)

func init() {
	flag.StringVar(&configPath, "config-path", pathToLogFile, "path to application config")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
