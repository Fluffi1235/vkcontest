package service

import (
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/dao"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"log"
	"strconv"
	"strings"
	"time"
)

func SaveInBd(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string) {
	dao.SaveInBd(date, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city)
}

func RegistrUser(messageinfo *model.Message) {
	dao.RegistrUser(messageinfo)
}

func SaveCity(city string, chatid int64) {
	dao.CityChange(city, chatid)
}

func ClearDb() {
	dao.ClearDb()
}

func GetNameCity(rowcity *sql.Rows) string {
	var city string
	for rowcity.Next() {
		if err := rowcity.Scan(&city); err != nil {
			log.Println(err)
		}
	}
	return city
}

func GetWeatherByNDays(limit string, chatid int64) []string {
	limitnum, _ := strconv.Atoi(limit)
	weather := model.Weather{}
	var counter int
	answer := make([]string, 1)
	rowcity := dao.GetUserCity(chatid)
	defer rowcity.Close()
	city := GetNameCity(rowcity)
	rows := dao.WeatherByNDays(limitnum, city)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&weather.Id, &weather.Day, &weather.TimesOfDay, &weather.Temp, &weather.Weather, &weather.Pressure,
			&weather.Humidity, &weather.Windspeed, &weather.Felt, &weather.City); err != nil {
			log.Println(err)
		}
		daystr := weather.Day.Format("2006-01-02")
		weather.TimesOfDay = strings.ToUpper(weather.TimesOfDay)
		if counter%4 == 0 {
			answer = append(answer, daystr+"\n"+weather.TimesOfDay+"\n"+weather.Temp+" "+weather.Weather+" Ощущается как "+
				weather.Felt+"\n"+"Давление = "+weather.Pressure+" мм рт.ст., Влажность = "+weather.Humidity+", Ветер = "+weather.Windspeed+"м/с\n\n")
		} else {
			answer = append(answer, weather.TimesOfDay+"\n"+weather.Temp+" "+weather.Weather+" Ощущается как "+
				weather.Felt+"\n"+"Давление = "+weather.Pressure+" мм рт.ст., Влажность = "+weather.Humidity+", Ветер = "+weather.Windspeed+"м/с\n\n")
		}
		counter++
	}
	return answer
}

func GetUserData(chatId int64) string {
	user := model.User{}
	rowData := dao.GetUserData(chatId)
	var answer string
	for rowData.Next() {
		if err := rowData.Scan(&user.ChatId, &user.UserName, &user.City, &user.FirstName, &user.LastName); err != nil {
			log.Println(err)
		}
		answer = "Ваши данные:\n" + "ChatId: " + strconv.Itoa(user.ChatId) + "\nUser Name: " + user.UserName + "\nИмя: " + user.FirstName + "\nФамилия: " + user.LastName +
			"\nГород: " + strings.Title(user.City)
	}
	return answer
}
