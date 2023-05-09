package repository

import (
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"time"
)

type UniversalRepo interface {
	SaveInBd(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string)
	RegistrUser(messageinfo *model.Message)
	CityChange(city string, chatid int64)
	ClearDb()
	GetUserData(chatId int64) *sql.Rows
	GetUserCity(chatId int64) *sql.Rows
	WeatherByNDays(limit int, city string) *sql.Rows
	GetNameCity(rowcity *sql.Rows) string
}
