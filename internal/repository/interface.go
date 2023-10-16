package repository

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/jmoiron/sqlx"
)

type UniversalRepo interface {
	UserRepo
	WeatherRepo
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UniversalRepo {
	return &Repository{
		db: db,
	}
}

type UserRepo interface {
	SetUser(messageInfo *model.MessageInfoText) error
	UpdateCityOfUser(city string, chatId int64) error
	GetUserData(chatId int64, platform string) (*model.User, error)
	GetCityOfUser(chatId int64) (string, error)
}

type WeatherRepo interface {
	SaveWeather(weather *model.Weather) error
	GetWeatherByNDays(limit int, city string) ([]model.Weather, error)
	ClearDb() error
}
