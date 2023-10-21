package services

import (
	"bytes"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parsers"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
	"text/template"
)

const (
	start = "Привет, я телеграм бот. Мои функции:\n" +
		"\t1)Вы можете посмотреть свои данные и в случае необходимости изменить город.\n" +
		"\t2)Посмотреть прогноз погоды в вашем городе до 10 дней. Не забудьте изменить город(город по умолчанию: Москва).\n" +
		"\tДоступные города: Москва, Санкт-Петербург, Новосибирск, Екатеренбург, Казань, Самара, Нижний Новгород, Ростов, Уфа, Железногорск (Курская обл.).\n" +
		"\t3)Использовать простой калькулятор. Функции калькулятора: сложение, вычитание, умножение, деление.\n" +
		"\t4)Использовать несколько беспланых API.\n" +
		"Функцианал API: узнать текущий курс BTC/USD, каллорийность некоторых фруктов.\n" +
		"Чтобы приступить к работе введите /info"
	tmplChangeCity string = "Ваш город сохранен  {{.}}\nТеперь вы можете смотреть погоду в вашем городе до 10 дней\nЧтобы посмотреть команды введите /info"
	tmplUserTG     string = "Ваши данные:\nChatId: {{.ChatId}}\nUser Name: @{{.UserName}}\nИмя: {{.FirstName}}\nФамилия: {{.LastName}}\nГород: {{.City}}"
	tmplUserVk     string = "Ваши данные:\nChatId: {{.ChatId}}\nИмя: {{.FirstName}}\nФамилия: {{.LastName}}\nГород: {{.City}}"
	tmplWeather    string = "{{.TimesOfDay}}\n{{.Temp}} {{.StateWeather}}, Ощущается как {{.Felt}}\nДавление = {{.Pressure}}мм рт.ст., Влажность = {{.Humidity}}, Ветер = {{.WindSpeed}}," +
		"м/с\n\n"
	tmplInfoFruit string = "Энергетическая ценность {{.Fruit}} на 100г:\nКалорийность {{.Calories}}кк\nПищевая ценность {{.Fruit}} \nЖиры {{.Fats}}г\nСахар {{.Sugar}}г\n" +
		"Углеводы {{.Carbohydrates}}г\nБелки {{.Protein}}г"
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
	return start
}

var changeCityTemplate = template.Must(template.New("changeCity").Parse(tmplChangeCity))
var userDataTgTemplate = template.Must(template.New("userDataTg").Parse(tmplUserTG))
var userDataVkTemplate = template.Must(template.New("userDataVk").Parse(tmplUserVk))
var weatherTemplate = template.Must(template.New("infoFruits").Parse(tmplWeather))
var fruitsTemplate = template.Must(template.New("infoFruit").Parse(tmplInfoFruit))

func (r *Repository) AnswerForCityChange(city string, chatId int64) (string, error) {
	err := r.repo.UpdateCityOfUser(city, chatId)
	if err != nil {
		return "", errors.Wrap(err, "[AnswerForCityChange]")
	}
	var buf bytes.Buffer
	err = changeCityTemplate.Execute(&buf, city)
	if err != nil {
		log.Println(err)
		return "", errors.Wrap(err, "[AnswerForCityChange]")
	}
	return buf.String(), nil
}

func (r *Repository) AnswerUserData(chatId int64, platform string) (string, error) {
	var buf bytes.Buffer
	var t *template.Template
	userData, err := r.repo.GetUserData(chatId, platform)
	if err != nil {
		return "", errors.Wrap(err, "[AnswerUserData]")
	}
	userData.City = strings.ToTitle(userData.City)
	switch userData.Platform {
	case "tg":
		t = userDataTgTemplate
	case "vk":
		t = userDataVkTemplate
	}
	err = t.Execute(&buf, userData)
	if err != nil {
		return "", errors.Wrap(err, "[AnswerUserData]")
	}
	return buf.String(), nil
}

func (r *Repository) AnswerWeatherByNDays(limit string, chatId int64) ([]string, error) {
	var counter int
	var answersForDay string
	var buf bytes.Buffer
	arrAnswer := make([]string, 1)
	amountDays, err := strconv.Atoi(limit)
	if err != nil {
		return nil, errors.Wrap(err, "[AnswerWeatherByNDays]")
	}
	userCity, err := r.repo.GetCityOfUser(chatId)
	if err != nil {
		return nil, errors.Wrap(err, "[AnswerWeatherByNDays]")
	}
	weather, err := r.repo.GetWeatherByNDays(amountDays, userCity)
	if err != nil {
		return nil, errors.Wrap(err, "[AnswerWeatherByNDays]")
	}
	for i := 0; i < len(weather); i++ {
		dayFormat := weather[i].Data.Format("2006-01-02")
		weather[i].TimesOfDay = strings.ToUpper(weather[i].TimesOfDay)
		if counter < 4 {
			if counter == 0 {
				answersForDay += dayFormat + "\n"
			}
			err = weatherTemplate.Execute(&buf, weather[i])
			if err != nil {
				return nil, errors.Wrap(err, "[AnswerWeatherByNDays]")
			}
			answersForDay += buf.String()
			counter++
		} else {
			arrAnswer = append(arrAnswer, answersForDay)
			counter = 0
		}
	}
	return arrAnswer, nil
}

func AnswerInfoFruits(fruit string, fruitInfo *parsers.Nutrit) (string, error) {
	infoFruit := model.InfoFruit{
		Fruit:         fruit,
		Calories:      strconv.FormatFloat(fruitInfo.Calories, 'f', 1, 64),
		Fats:          strconv.FormatFloat(fruitInfo.Fats, 'f', 1, 64),
		Sugar:         strconv.FormatFloat(fruitInfo.Sugar, 'f', 1, 64),
		Carbohydrates: strconv.FormatFloat(fruitInfo.Carbohydrates, 'f', 1, 64),
		Protein:       strconv.FormatFloat(fruitInfo.Protein, 'f', 1, 64),
	}
	var buf bytes.Buffer
	err := fruitsTemplate.Execute(&buf, infoFruit)
	if err != nil {
		return "", errors.Wrap(err, "[AnswerInfoFruits]")
	}
	return buf.String(), nil
}
