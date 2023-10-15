package repository

import (
	"database/sql"
	"log"
	"time"
)

func (r Repository) SaveWeather(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string) {
	_, err := r.db.Exec("INSERT INTO weather_forecast(day, timesofday, temp, weather, pressure, humidity, windSpeed, felt, city) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		date, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city)
	if err != nil {
		log.Println("Error inserting into dao SaveWeather")
	}
}

func (r Repository) ClearDb() {
	_, err := r.db.Exec("DELETE FROM weather_forecast")
	if err != nil {
		log.Println("Error delete table weather_forecast")
	}
}

func (r Repository) WeatherByNDays(limit int, city string) *sql.Rows {
	rows, err := r.db.Query("SELECT day, timesofday, temp, weather, pressure, humidity, windspeed, felt, city FROM weather_forecast where city = $1 ORDER BY id Limit $2",
		city, limit*4)
	if err != nil {
		log.Println("Error select data WeatherByNDays")
	}
	return rows
}
