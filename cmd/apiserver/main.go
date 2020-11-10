package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/elmira-aliyeva/go-rest-api/internal/app/apiserver"
)

var configPath string

// set the config-path flag, set default value, description
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	// parses flags from args
	flag.Parse()

	// NewConfig returns an instance of Config struct,
	// with default configs: port - 8080, log level - debug (databaseURL, sessionKey are empty)
	config := apiserver.NewConfig()

	// take config from toml file and write values to config struct
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Start checks the db, starts listening on the given port
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
