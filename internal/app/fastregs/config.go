package fastregs

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	databaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    "8080",
		LogLevel:    "debug",
		databaseURL: "host=localhost dbname=fastregs sslmode=disable",
	}
}
