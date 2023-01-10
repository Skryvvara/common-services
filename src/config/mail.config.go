package config

type MailConfig struct {
	Host     string `toml:"host" env:"MAIL_HOST"`
	Port     int    `toml:"port" env:"MAIL_PORT" default:"24"`
	Secure   bool   `toml:"secure" env:"MAIL_SECURE" default:"false"`
	User     string `toml:"user" env:"MAIL_USER"`
	Password string `toml:"password" env:"MAIL_PASSWORD"`
	To       string `toml:"to" env:"MAIL_TO"`
}
