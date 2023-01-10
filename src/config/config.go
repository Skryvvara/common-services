package config

import (
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
	"github.com/creasty/defaults"
)

type Config struct {
	Server ServerConfig `toml:"server"`
	Log    LogConfig    `toml:"log"`
	Mail   MailConfig   `toml:"mail"`
}

var App Config

func getEnvValue(key, fallback string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}
	return fallback
}

func Initialize() {
	log.Println("Starting config initialization")
	defer log.Println("Config successfully loaded")

	configPath := path.Clean(getEnvValue("CONFIG_PATH", path.Join("data", "config.toml")))
	if _, err := os.Stat(configPath); err != nil {
		if _, err := os.Stat("config.toml"); err != nil {
			log.Panic(err)
		}
		configPath = "config.toml"
	}

	bytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Panic(err)
	}

	_, err = toml.Decode(string(bytes), &App)
	if err != nil {
		log.Panic(err)
	}

	if err := env.Parse(&App); err != nil {
		log.Println(err)
	}

	App.Server.ConfigPath = configPath
	defaults.Set(&App)
}
