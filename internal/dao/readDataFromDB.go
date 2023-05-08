package dao

import (
	"database/sql"
	"log"
	"telegram_bot/resources"
)

func GetUserCity(chatid int64) *sql.Rows {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
	}
	defer db.Close()
	row, err := db.Query("SELECT city FROM users where chatid = $1", chatid)
	if err != nil {
		log.Println(err)
	}
	return row
}

func WeatherByDate(date, city string) *sql.Rows {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM weather_forecast where day = $1 and city = $2 ORDER BY id ", date, city)
	if err != nil {
		log.Println(err)
	}
	return rows
}

func WeatherByInterval(firstdate, lastdate, city string) *sql.Rows {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM weather_forecast where day >= $1 and day <= $2 and city = $3 ORDER BY id ", firstdate, lastdate, city)
	if err != nil {
		log.Println(err)
	}
	return rows
}

func WeatherByNDays(limit int, city string) *sql.Rows {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM weather_forecast where city = $1 ORDER BY id Limit $2", city, limit*4)
	if err != nil {
		log.Println(err)
	}
	return rows
}
