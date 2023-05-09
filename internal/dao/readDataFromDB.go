package dao

import (
	"database/sql"
	"github.com/Fluffi1235/vkcontest/resources"
	"log"
)

func GetUserData(chatId int64) *sql.Rows {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
	}
	defer db.Close()
	row, err := db.Query("SELECT * FROM users where chatid = $1", chatId)
	if err != nil {
		log.Println(err)
	}
	return row
}

func GetUserCity(chatId int64) *sql.Rows {
	db, err := sql.Open("postgres", resources.DbConnect)
	if err != nil {
		log.Println("Error connecting to dao")
	}
	defer db.Close()
	row, err := db.Query("SELECT city FROM users where chatid = $1", chatId)
	if err != nil {
		log.Println(err)
	}
	return row
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
