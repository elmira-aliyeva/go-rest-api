package apiserver

import "github.com/elmira-aliyeva/go-rest-api/internal/store"

//Config ...
type Config struct {
	BindAddr string        `toml:"bind_addr"`
	LogLevel string        `toml:"log_level"`
	Store    *store.Config // DataBaseURL string `toml:"database_url"`
}

// NewConfig returns an instance of Config struct,
// with default configs: port - 8080, log level - debug, store - (no val)
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(), // DataBaseURL string `toml:"database_url"`
	}
}
