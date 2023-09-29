package repository

import (
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"time"
)

type UniversalRepo interface {
	UserRepo
	WheatherRepo
}

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) UniversalRepo {
	return &Repository{
		db: db,
	}
}

type UserRepo interface {
	RegistrUser(messageinfo *model.MessageInfoText)
	CityChange(city string, chatid int64)
	GetUserData(chatId int64, platform string) *sql.Rows
	GetCityOfUser(chatId int64) *sql.Rows
}

type WheatherRepo interface {
	SaveWheather(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string)
	WeatherByNDays(limit int, city string) *sql.Rows
	ClearDb()
}
