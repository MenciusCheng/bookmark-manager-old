package conf

import (
	"github.com/MenciusCheng/superman/util/dragons"
)

type Config struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
}

func Init() (*Config, error) {
	// parse Config from config file
	cfg := &Config{}
	err := dragons.Scan(cfg)
	return cfg, err
}
