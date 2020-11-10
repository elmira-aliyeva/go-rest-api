package apiserver

//Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DataBaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// NewConfig returns an instance of Config struct,
// with default configs: port - 8080, log level - debug (databaseURL, sessionKey are empty)
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
