package model

import "time"

type Weather struct {
	Id         int
	Day        time.Time
	TimesOfDay string
	Temp       string
	Weather    string
	Pressure   string
	Humidity   string
	Windspeed  string
	Felt       string
	City       string
}
