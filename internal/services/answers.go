package services

import (
	"fmt"
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
	answer := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		"Привет, я телеграм бот. Мои функции:",
		"\t1)Вы можете посмотреть свои данные и в случае необходимости изменить город.",
		"\t2)Посмотреть прогноз погоды в вашем городе до 10 дней. Не забудьте изменить город(город по умолчанию: Москва).",
		"\tДоступные города: Москва, Санкт-Петербург, Новосибирск, Екатеренбург, Казань, Самара, Нижний Новгород, Ростов, Уфа, Железногорск (Курская обл.).",
		"\t3)Использовать простой калькулятор. Функции калькулятора: сложение, вычитание, умножение, деление.",
		"\t4)Использовать несколько беспланых API.",
		"Функцианал API: узнать текущий курс BTC/USD, каллорийность некоторых фруктов.",
		"Чтобы приступить к работе введите /info")
	return answer
}

func (r *Repository) AnswerForCityChange(city string, chatId int64) string {
	r.repo.CityChange(city, chatId)
	answer := fmt.Sprintf("%s%s\n%s\n%s",
		"Ваш город сохранен ", strings.Title(city),
		"Теперь вы можете смотреть погоду в вашем городе до 10 дней",
		"Чтобы посмотреть команды введите /info")
	return answer
}

func (r *Repository) AnswerUserData(chatId int64, platform string) string {
	user := model.User{}
	rowData := r.repo.GetUserData(chatId, platform)
	var answer string
	for rowData.Next() {
		if err := rowData.Scan(&user.ChatId, &user.UserName, &user.City, &user.FirstName, &user.LastName, &user.Platform); err != nil {
			log.Println(err)
		}
		if user.Platform == "tg" {
			answer = fmt.Sprintf("%s\n%s%d\n%s%s\n%s%s\n%s%s\n%s%s",
				"Ваши данные:",
				"ChatId: ", user.ChatId,
				"User Name: @", user.UserName,
				"Имя: ", user.FirstName,
				"Фамилия: ", user.LastName,
				"Город: ", strings.Title(user.City))
		}
		if user.Platform == "vk" {
			answer = fmt.Sprintf("%s\n%s%d\n%s%s\n%s%s\n%s%s",
				"Ваши данные:",
				"ChatId: ", user.ChatId,
				"Имя: ", user.FirstName,
				"Фамилия: ", user.LastName,
				"Город: ", strings.Title(user.City))
		}
	}
	return answer
}

func (r *Repository) AnswerWeatherByNDays(limit string, chatid int64) []string {
	amountDays, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
	}
	var counter int
	answer := make([]string, 1)
	userCity := r.repo.GetCityOfUser(chatid)
	defer userCity.Close()
	var city string
	for userCity.Next() {
		if err = userCity.Scan(&city); err != nil {
			log.Println(err)
		}
	}
	rows := r.repo.WeatherByNDays(amountDays, city)
	defer rows.Close()
	weather := model.Weather{}
	for rows.Next() {
		if err = rows.Scan(&weather.Id, &weather.Day, &weather.TimesOfDay, &weather.Temp, &weather.Weather, &weather.Pressure,
			&weather.Humidity, &weather.WindSpeed, &weather.Felt, &weather.City); err != nil {
			log.Println(err)
		}
		daystr := weather.Day.Format("2006-01-02")
		weather.TimesOfDay = strings.ToUpper(weather.TimesOfDay)
		if counter%4 == 0 {
			answer = append(answer, strings.Join([]string{daystr, "\n", weather.TimesOfDay, "\n", weather.Temp, " ", weather.Weather, " Ощущается как ",
				weather.Felt, "\n", "Давление = ", weather.Pressure, " мм рт.ст., Влажность = ", weather.Humidity, ", Ветер = ", weather.WindSpeed, "м/с\n\n",
			}, ""))
		} else {
			answer = append(answer, strings.Join([]string{
				weather.TimesOfDay, "\n", weather.Temp, " ", weather.Weather, " Ощущается как ",
				weather.Felt, "\n", "Давление = ", weather.Pressure, " мм рт.ст., Влажность = ", weather.Humidity, ", Ветер = ", weather.WindSpeed, "м/с\n\n",
			}, ""))
		}
		counter++
	}
	return answer
}
