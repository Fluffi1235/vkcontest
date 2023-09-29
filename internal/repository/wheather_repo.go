package repository

import (
	"database/sql"
	"log"
	"time"
)

func (r Repository) SaveWheather(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string) {
	_, err := r.db.Exec("INSERT INTO weather_forecast(day, timesofday, temp, weather, pressure, humidity, windspeed, felt, city) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		date, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city)
	if err != nil {
		log.Println(err)
	}
}

func (r Repository) ClearDb() {
	_, err := r.db.Exec("DELETE FROM weather_forecast")
	if err != nil {
		log.Fatal(err)
	}
}

func (r Repository) WeatherByNDays(limit int, city string) *sql.Rows {
	rows, err := r.db.Query("SELECT * FROM weather_forecast where city = $1 ORDER BY id Limit $2", city, limit*4)
	if err != nil {
		log.Println(err)
	}
	return rows
}
