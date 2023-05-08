package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"telegram_bot/internal/model"
	"telegram_bot/resources"
	"time"
)

func SaveInBd(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string) {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO weather_forecast(day, timesofday, temp, weather, pressure, humidity, windspeed, felt, city) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		date, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city)
	if err != nil {
		log.Println("Error inserting into save in Bd")
		return
	}
}

func RegistrUser(messageinfo *model.Message) {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT *From Users where chatid = $1", messageinfo.ChatID)
	if rows != nil {
		return
	}
	_, err = db.Exec("INSERT INTO Users(chatid, username, city, first_name, last_name) VALUES($1, $2, $3, $4, $5)",
		messageinfo.ChatID, messageinfo.Username, "москва", messageinfo.FirstName, messageinfo.LastName)
	if err != nil {
		log.Println("Error inserting into dao RegistrUser")
		return
	}
}

func SaveCity(city string, chatid int64) {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE users set city = $1 where chatid = $2", city, chatid)
	if err != nil {
		log.Println("Error inserting into dao SaveCity")
		return
	}
}

func ClearDb() {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM weather_forecast")
	if err != nil {
		log.Println(err)
	}
}
