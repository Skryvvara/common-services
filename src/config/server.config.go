package config

type ServerConfig struct {
	Environment     string `toml:"environment" env:"SERVER_ENVIRONMENT" default:"PRODUCTION"`
	Port            string `toml:"port" env:"SERVER_PORT" default:"3000"`
	TimeZone        string `toml:"timezone" env:"SERVER_TIMEZONE" default:"Europe/Berlin"`
	PoweredByHeader string `toml:"powered_by_header" env:"SERVER_POWERED_BY_HEADER" default:"calyx"`
	ConfigPath      string
}
