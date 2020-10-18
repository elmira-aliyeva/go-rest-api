package apiserver

//Config ...
type Config struct {
	BindAddr string `tool:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
