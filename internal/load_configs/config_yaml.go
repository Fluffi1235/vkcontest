package load_configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	ConnectDb         string `yaml:"connectdb"`
	TgToken           string `yaml:"tgtoken"`
	VkToken           string `yaml:"vktoken"`
	WeatherUpdateInfo int    `yaml:"weatherUpdateInfo"`
	FruityViceLink    string `yaml:"fruityVice"`
	CoinDeskLink      string `yaml:"coinDesk"`
}

func LoadConfigFromYaml() (*Config, error) {
	var conf Config
	f, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
