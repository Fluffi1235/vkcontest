package buttons

import (
	"github.com/Fluffi1235/vkcontest/internal/load_configs"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parsers"
	"github.com/Fluffi1235/vkcontest/internal/services"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func DataUser(msg string, chatId int64, servise *services.Repository, platform string) (bool, string) {
	var answer string

	checkPrefData, err := regexp.MatchString("показать мои данные", msg)
	if err != nil {
		log.Println(err)
	}
	if checkPrefData {
		answer = servise.AnswerUserData(chatId, platform)
		return true, answer
	}
	return false, answer
}

func CheckCity(msg string, chatId int64, servise *services.Repository) (bool, string) {
	cities := model.Cities()
	var city string
	var answer string
	for k := range cities {
		checkPrefDay, err := regexp.MatchString("city .+", msg)
		if err != nil {
			log.Println(err)
		}
		if checkPrefDay {
			city = strings.Split(msg, "city ")[1]
			if city == k {
				answer = servise.AnswerForCityChange(city, chatId)
				return true, answer
			} else {
				answer = "Информации о погоде об этом городе нет"
			}
		}
	}
	return false, answer
}

func WeatherOnNDays(msg string, chatId int64, servise *services.Repository) (bool, []string) {
	arrAnswer := make([]string, 0)
	checkPrefInterval, err := regexp.MatchString("погода \\d{1,10}", msg)
	if err != nil {
		log.Println(err)
	}
	if checkPrefInterval {
		msgSplit := strings.Split(msg, " ")
		arrAnswer = servise.AnswerWeatherByNDays(msgSplit[1], chatId)
		return true, arrAnswer
	}
	return false, arrAnswer
}

func Calculator(msg string, chatId int64, person map[int64]rune) bool {
	checkPrefCalc, err := regexp.MatchString("calc [-+*/]", msg)
	if err != nil {
		log.Println(err)
	}
	if checkPrefCalc {
		symbol := msg[len(msg)-1]
		person[chatId] = rune(symbol)
		return true
	}
	return false
}

func IsTwoNumbers(msg string, chatId int64, person map[int64]rune) string {
	var res float64
	var answer string
	checkPrefTwoNumbers, err := regexp.MatchString("[-]?\\d+[.,]?\\d* [-]?\\d+[.,]?\\d*", msg)
	if err != nil {
		log.Println(err)
	}
	if checkPrefTwoNumbers {
		msgSplit := strings.Split(msg, " ")
		msgSplit[0] = strings.ReplaceAll(msgSplit[0], ",", ".")
		msgSplit[1] = strings.ReplaceAll(msgSplit[1], ",", ".")
		num1, err := strconv.ParseFloat(msgSplit[0], 64)
		if err != nil {
			log.Println(err)
		}
		num2, err := strconv.ParseFloat(msgSplit[1], 64)
		if err != nil {
			log.Println(err)
		}
		switch person[chatId] {
		case '+':
			res = num1 + num2
		case '-':
			res = num1 - num2
		case '*':
			res = num1 * num2
		case '/':
			if num2 == 0 {
				return "На 0 делить нельзя"
			}
			res = num1 / num2
		default:
			return "Не выбрали операцию в калькуляторе"
		}
		delete(person, chatId)
		answer = strconv.FormatFloat(res, 'f', 1, 64)
	}
	return answer
}

func InfoFruits(msg string, config *load_configs.Config) (bool, string) {
	checkPrefFruit, err := regexp.MatchString("apifruit .+", msg)
	if err != nil {
		log.Println(err)
	}
	if checkPrefFruit {
		fruit := strings.Split(msg, " ")
		fruitInfo := parsers.ParseFruit(fruit[1], config)
		answer := services.AnswerInfoFruits(fruit[2], &fruitInfo)
		return true, answer
	}
	return false, ""
}

func Btc(msg string, config *load_configs.Config) (bool, string) {
	var answer string
	checkPrefBtc, err := regexp.MatchString("btc", msg)
	if err != nil {
		log.Println(err)
	}
	if checkPrefBtc {
		answer = parsers.ParseBtc(config)
		return true, answer
	}
	return false, answer
}
