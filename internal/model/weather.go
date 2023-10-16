package model

import "time"

type Weather struct {
	Id           int
	Data         time.Time
	TimesOfDay   string
	Temp         string
	StateWeather string
	Pressure     string
	Humidity     string
	WindSpeed    string
	Felt         string
	City         string
}
