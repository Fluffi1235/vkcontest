package repository

import (
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"log"
	"time"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) UniversalRepo {
	return &Repository{
		db: db,
	}
}

func (r Repository) SaveInBd(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string) {
	_, err := r.db.Exec("INSERT INTO weather_forecast(day, timesofday, temp, weather, pressure, humidity, windspeed, felt, city) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		date, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city)
	if err != nil {
		log.Println("Error inserting into save in Bd")
		return
	}
}

func (r Repository) RegistrUser(messageinfo *model.Message) {
	rows, err := r.db.Query("SELECT *From Users where chatid = $1", messageinfo.ChatID)
	if rows != nil {
		return
	}
	_, err = r.db.Exec("INSERT INTO Users(chatid, username, city, first_name, last_name) VALUES($1, $2, $3, $4, $5)",
		messageinfo.ChatID, messageinfo.Username, "москва", messageinfo.FirstName, messageinfo.LastName)
	if err != nil {
		log.Println("Error inserting into dao RegistrUser")
		return
	}
}

func (r Repository) CityChange(city string, chatid int64) {
	_, err := r.db.Exec("UPDATE users set city = $1 where chatid = $2", city, chatid)
	if err != nil {
		log.Println("Error inserting into dao CityChange")
		return
	}
}

func (r Repository) ClearDb() {
	_, err := r.db.Exec("DELETE FROM weather_forecast")
	if err != nil {
		log.Println(err)
	}
}

func (r Repository) GetUserData(chatId int64) *sql.Rows {
	row, err := r.db.Query("SELECT * FROM users where chatid = $1", chatId)
	if err != nil {
		log.Println(err)
	}
	return row
}

func (r Repository) GetUserCity(chatId int64) *sql.Rows {
	row, err := r.db.Query("SELECT city FROM users where chatid = $1", chatId)
	if err != nil {
		log.Println(err)
	}
	return row
}

func (r Repository) WeatherByNDays(limit int, city string) *sql.Rows {
	rows, err := r.db.Query("SELECT * FROM weather_forecast where city = $1 ORDER BY id Limit $2", city, limit*4)
	if err != nil {
		log.Println(err)
	}
	return rows
}

func (r Repository) GetNameCity(rowcity *sql.Rows) string {
	var city string
	for rowcity.Next() {
		if err := rowcity.Scan(&city); err != nil {
			log.Println(err)
		}
	}
	return city
}
