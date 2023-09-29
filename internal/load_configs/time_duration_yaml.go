package load_configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type TimeDuration struct {
	WeatherUpdateInfo int `yaml:"WeatherUpdateInfo"`
}

func LoadConfigTimeDuration() (*TimeDuration, error) {
	var t TimeDuration
	f, err := ioutil.ReadFile("./config/TimeDuration.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(f, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
