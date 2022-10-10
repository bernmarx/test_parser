package config

type Config struct {
	Server struct {
		BindAddr string `yaml:"bind_addr"`
		LogLevel string `yaml:"log_level"`
		BaseLink string `yaml:"link"`
	} `yaml:"server"`
	Store struct {
		DatabaseURL string `yaml:"database_url"`
	} `yaml:"store"`
}
