package services

import (
	"bytes"
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parsers"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"log"
	"strconv"
	"strings"
	"text/template"
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
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		"Привет, я телеграм бот. Мои функции:",
		"\t1)Вы можете посмотреть свои данные и в случае необходимости изменить город.",
		"\t2)Посмотреть прогноз погоды в вашем городе до 10 дней. Не забудьте изменить город(город по умолчанию: Москва).",
		"\tДоступные города: Москва, Санкт-Петербург, Новосибирск, Екатеренбург, Казань, Самара, Нижний Новгород, Ростов, Уфа, Железногорск (Курская обл.).",
		"\t3)Использовать простой калькулятор. Функции калькулятора: сложение, вычитание, умножение, деление.",
		"\t4)Использовать несколько беспланых API.",
		"Функцианал API: узнать текущий курс BTC/USD, каллорийность некоторых фруктов.",
		"Чтобы приступить к работе введите /info")
}

func (r *Repository) AnswerForCityChange(city string, chatId int64) string {
	r.repo.CityChange(city, chatId)
	tmpl := "Ваш город сохранен  {{.}}\nТеперь вы можете смотреть погоду в вашем городе до 10 дней\nЧтобы посмотреть команды введите /info"
	t := template.Must(template.New("changeCity").Parse(tmpl))
	var buf bytes.Buffer
	err := t.Execute(&buf, city)
	if err != nil {
		fmt.Println(err)
	}
	return buf.String()
}

func (r *Repository) AnswerUserData(chatId int64, platform string) string {
	user := model.User{}
	var buf bytes.Buffer
	var t *template.Template
	rowData := r.repo.GetUserData(chatId, platform)
	for rowData.Next() {
		if err := rowData.Scan(&user.ChatId, &user.UserName, &user.City, &user.FirstName, &user.LastName, &user.Platform); err != nil {
			log.Println(err)
		}
		user.City = strings.Title(user.City)
		if user.Platform == "tg" {
			tmplTG := "Ваши данные:\nChatId: {{.ChatId}}\nUser Name: @{{.UserName}}\nИмя: {{.FirstName}}\nФамилия: {{.LastName}}\nГород: {{.City}}"
			t = template.Must(template.New("userDataTg").Parse(tmplTG))
		}
		if user.Platform == "vk" {
			tmplTG := "Ваши данные:\nChatId: {{.ChatId}}\nИмя: {{.FirstName}}\nФамилия: {{.LastName}}\nГород: {{.City}}"
			t = template.Must(template.New("userDataVk").Parse(tmplTG))
		}
	}
	err := t.Execute(&buf, user)
	if err != nil {
		fmt.Println(err)
	}
	return buf.String()
}

func (r *Repository) AnswerWeatherByNDays(limit string, chatid int64) []string {
	amountDays, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
	}
	var counter int
	arrAnswer := make([]string, 1)
	answersForDay := ""
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
		dayFormat := weather.Day.Format("2006-01-02")
		weather.TimesOfDay = strings.ToUpper(weather.TimesOfDay)
		if counter < 4 {
			if counter == 0 {
				answersForDay += dayFormat + "\n"
			}
			answersForDay += strings.Join([]string{
				weather.TimesOfDay, "\n", weather.Temp, " ", weather.Weather, " Ощущается как ",
				weather.Felt, "\n", "Давление = ", weather.Pressure, " мм рт.ст., Влажность = ", weather.Humidity, ", Ветер = ", weather.WindSpeed, "м/с\n\n",
			}, "")
			counter++
		} else {
			arrAnswer = append(arrAnswer, answersForDay)
			counter = 0
		}
	}
	return arrAnswer
}

func AnswerInfoFruits(fruit string, fruitInfo *parsers.Nutrit) string {
	infoFruit := struct {
		Fruit         string
		Calories      string
		Fats          string
		Sugar         string
		Carbohydrates string
		Protein       string
	}{fruit,
		strconv.FormatFloat(fruitInfo.Calories, 'f', 1, 64),
		strconv.FormatFloat(fruitInfo.Fats, 'f', 1, 64),
		strconv.FormatFloat(fruitInfo.Sugar, 'f', 1, 64),
		strconv.FormatFloat(fruitInfo.Carbohydrates, 'f', 1, 64),
		strconv.FormatFloat(fruitInfo.Protein, 'f', 1, 64),
	}
	tmpl := "Энергетическая ценность {{.Fruit}} на 100г:\nКалорийность {{.Calories}}кк\nПищевая ценность {{.Fruit}} \nЖиры {{.Fats}}г\nСахар {{.Sugar}}г\n" +
		"Углеводы {{.Carbohydrates}}г\nБелки {{.Protein}}г"
	t := template.Must(template.New("changeCity").Parse(tmpl))
	var buf bytes.Buffer
	err := t.Execute(&buf, infoFruit)
	if err != nil {
		fmt.Println(err)
	}
	return buf.String()
}
