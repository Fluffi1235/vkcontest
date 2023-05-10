package service

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"log"
	"strconv"
	"strings"
)

type Repository struct {
	repo repository.UniversalRepo
}

func New(repo repository.UniversalRepo) *Repository {
	return &Repository{
		repo: repo,
	}
}

func (r *Repository) AnswerStart() string {
	answer := "Привет, я телеграм бот. Мои функции:\n" +
		"1)Вы можете посмотреть свои данные и в случае необходимости изменить город.\n" +
		"2)Посмотреть прогноз погоды в вашем городе до 10 дней. Не забудьте изменить город(город по умолчанию: Москва).\n" +
		"Доступные города: Москва, Санкт-Петербург, Новосибирск, Екатеренбург, Казань, Самара, Нижний Новгород, Ростов, Уфа, Железногорск (Курская обл.).\n" +
		"3)Использовать простой калькулятор.\n" +
		"Функции калькулятора: сложение, вычитание, умножение, деление.\n" +
		"4)Использовать несколько беспланых API.\n" +
		"Функцианал API: узнать текущий курс BTC/USD, каллорийность некоторых фруктов.\n" +
		"Чтобы приступить к работе введите /info"

	return answer
}

func (r *Repository) AnswerForCityChange(city string) string {
	answer := "Ваш город сохранен " + strings.Title(city) +
		"\nТеперь вы можете смотреть погоду в вашем городе до 10 дней\n" +
		"Чтобы посмотреть команды введите /info"
	return answer
}

func (r *Repository) GetUserCity(chatId int64, platform string) string {
	user := model.User{}
	rowData := r.repo.GetUserData(chatId, platform)
	var answer string
	for rowData.Next() {
		if err := rowData.Scan(&user.ChatId, &user.UserName, &user.City, &user.FirstName, &user.LastName, &user.Platform); err != nil {
			log.Println(err)
		}
		if user.Platform == "tg" {
			answer = "Ваши данные:\n" + "ChatId: " + strconv.Itoa(user.ChatId) + "\nUser Name: @" + user.UserName + "\nИмя: " + user.FirstName + "\nФамилия: " + user.LastName +
				"\nГород: " + strings.Title(user.City)
		}
		if user.Platform == "vk" {
			answer = "Ваши данные:\n" + "ChatId: " + strconv.Itoa(user.ChatId) + "\nUser Name: @" + user.UserName + "\nИмя: " + user.FirstName + "\nФамилия: " + user.LastName +
				"\nГород: " + strings.Title(user.City)
		}
	}
	return answer
}

func (r *Repository) GetWeatherByNDays(limit string, chatid int64) []string {
	limitnum, _ := strconv.Atoi(limit)
	weather := model.Weather{}
	var counter int
	answer := make([]string, 1)
	rowcity := r.repo.GetUserCity(chatid)
	defer rowcity.Close()
	city := r.repo.GetNameCity(rowcity)
	rows := r.repo.WeatherByNDays(limitnum, city)
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

func (r *Repository) SaveCity(city string, chatid int64) {
	r.repo.CityChange(city, chatid)
}
