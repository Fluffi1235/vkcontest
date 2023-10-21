package repository

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/pkg/errors"
)

func (r Repository) SaveWeather(weather *model.Weather) error {
	query := "INSERT INTO weather_forecast(day, timesofday, temp, weather, pressure, humidity, windSpeed, felt, city) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := r.db.Exec(query,
		weather.Data, weather.TimesOfDay, weather.Temp, weather.StateWeather, weather.Pressure, weather.Humidity, weather.WindSpeed, weather.Felt, weather.City)
	if err != nil {
		return errors.Wrap(err, "[Dao SaveWeather]")
	}
	return nil
}

func (r Repository) ClearDb() error {
	_, err := r.db.Exec("TRUNCATE FROM weather_forecast")
	if err != nil {
		return errors.Wrap(err, "[Dao ClearDb]")
	}
	return nil
}

func (r Repository) GetWeatherByNDays(limit int, city string) ([]model.Weather, error) {
	weather := []model.Weather{}
	query := "SELECT day, timesofday, temp, weather, pressure, humidity, windspeed, felt, city FROM weather_forecast where city = $1 ORDER BY id Limit $2"
	err := r.db.Select(&weather, query, city, limit*4)
	if err != nil {
		return nil, errors.Wrap(err, "[Dao GetWeatherByNDays]")
	}
	return weather, nil
}
