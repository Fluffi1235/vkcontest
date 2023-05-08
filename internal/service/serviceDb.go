package service

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"telegram_bot/internal/dao"
	"telegram_bot/internal/model"
	"time"
)

func SaveInBd(date time.Time, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city string) {
	dao.SaveInBd(date, timesOfDay, temp, weather, pressure, humidity, windspeed, felt, city)
}

func RegistrUser(messageinfo *model.Message) {
	dao.RegistrUser(messageinfo)
}

func SaveCity(city string, chatid int64) {
	dao.SaveCity(city, chatid)
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

func GetWeatherByDate(date string, chatid int64) string {
	weather := model.Weather{}
	var answer string
	rowcity := dao.GetUserCity(chatid)
	defer rowcity.Close()
	city := GetNameCity(rowcity)
	rowsweather := dao.WeatherByDate(date, city)
	defer rowsweather.Close()
	for rowsweather.Next() {
		if err := rowsweather.Scan(&weather.Id, &weather.Day, &weather.TimesOfDay, &weather.Temp, &weather.Weather,
			&weather.Pressure, &weather.Humidity, &weather.Windspeed, &weather.Felt, &weather.City); err != nil {
			log.Println(err)
		}
		weather.TimesOfDay = strings.ToUpper(weather.TimesOfDay)
		answer = answer + weather.TimesOfDay + "\n" + weather.Temp + ", " + weather.Weather + " Ощущается как " +
			weather.Felt + "\n" + "Давление = " + weather.Pressure + " мм рт.ст., Влажность = " + weather.Humidity +
			", Ветер = " + weather.Windspeed + "м/с\n\n"
	}
	return answer
}

func GetWeatherByInterval(firstdate, lastdate string, chatid int64) []string {
	weather := model.Weather{}
	answer := make([]string, 1)
	var counter int
	rowcity := dao.GetUserCity(chatid)
	defer rowcity.Close()
	city := GetNameCity(rowcity)
	rows := dao.WeatherByInterval(firstdate, lastdate, city)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&weather.Id, &weather.Day, &weather.TimesOfDay, &weather.Temp, &weather.Weather,
			&weather.Pressure, &weather.Humidity, &weather.Windspeed, &weather.Felt, &weather.City); err != nil {
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
