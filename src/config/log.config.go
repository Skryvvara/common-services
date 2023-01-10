package config

type LogConfig struct {
	LogDir     string `toml:"log_dir" env:"LOG_DIR" default:"log"`
	LogFile    string `toml:"log_file" env:"LOG_FILE" default:"app.log"`
	MaxSize    int    `toml:"max_size_in_mb" env:"LOG_MAX_SIZE_IN_MB" default:"1"`
	MaxBackups int    `toml:"max_backups" env:"LOG_MAX_BACKUPS" default:"3"`
	MaxAge     int    `toml:"max_age" env:"LOG_MAX_AGE" default:"28"`
	Color      bool   `toml:"color" env:"LOG_COLOR" default:"false"`
}
