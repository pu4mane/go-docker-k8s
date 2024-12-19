package config

import (
	"errors"
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	ErrEmptyPath   = errors.New("config path is empty")
	ErrReadCfg     = errors.New("failed to read config file")
	ErrCfgNotExist = errors.New("config file does not exist")
)

type Config struct {
	Http HttpConfig `yaml:"http"`
}

type HttpConfig struct {
	Port string `yaml:"port"`
}

func MustLoad() (*Config, error) {
	cfgPath := fetchCfgPath()
	if cfgPath == "" {
		return nil, ErrEmptyPath
	}

	_, err := os.Stat(cfgPath)
	if os.IsNotExist(err) {
		return nil, ErrCfgNotExist
	}

	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		return nil, ErrReadCfg
	}

	return &cfg, nil
}

func fetchCfgPath() string {
	var cfgPath string

	flag.StringVar(&cfgPath, "config", "", "path to config file")

	flag.Parse()
	if cfgPath == "" {
		cfgPath = os.Getenv("CONFIG_PATH")
	}

	return cfgPath
}
