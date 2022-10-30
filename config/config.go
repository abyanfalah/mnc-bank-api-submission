package config

import (
	"os"
)

type ApiConfig struct {
	Host string
	Port string
}

type Config struct {
	ApiConfig
}

func (c *Config) readConfig() {
	c.ApiConfig = ApiConfig{
		Host: os.Getenv("API_HOST"),
		Port: os.Getenv("API_PORT"),
	}

}

func NewConfig() Config {
	conf := new(Config)
	conf.readConfig()
	return *conf
}
