package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	ConnectDb         string        `yaml:"connectdb"`
	TgToken           string        `yaml:"tgtoken"`
	VkToken           string        `yaml:"vktoken"`
	WeatherUpdateInfo time.Duration `yaml:"weatherUpdateInfo"`
	FruityViceLink    string        `yaml:"fruityVice"`
	CoinDeskLink      string        `yaml:"coinDesk"`
	ClientTimeout     time.Duration `yaml:"clientTimeout"`
	ContextTimeout    time.Duration `yaml:"contextTimeout"`
}

func LoadConfigFromYaml() (*Config, error) {
	var cfg *Config
	f, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
