package bot

import (
	"regexp"
	"strconv"
	"strings"
	"vkcontest/internal/model"
	"vkcontest/internal/parse"
	"vkcontest/internal/service"
)

func DataUser(msg string, chatId int64) (bool, string) {
	var answer string
	checkprefdata, _ := regexp.MatchString("Показать мои данные", msg)
	if checkprefdata {
		answer = service.GetUserData(chatId)
		return true, answer
	}
	return false, answer
}

func CheckCity(msg string, chatId int64) (bool, string) {
	citys := model.City()
	var city string
	var answer string
	for k, _ := range citys {
		checkprefday, _ := regexp.MatchString("city .+", msg)
		if checkprefday {
			city = strings.Split(msg, "city ")[1]
			if city == k {
				answer = service.AnswerForCityChange(city)
				service.SaveCity(k, chatId)
				return true, answer
			} else {
				answer = "Информации о погоде об этом городе нет"
			}
		}
	}
	return false, answer
}

func CheckNdays(msg string, chatId int64) (bool, []string) {
	answermas := make([]string, 0)
	checkprefinterval, _ := regexp.MatchString("Погода \\d{1,10}", msg)
	if checkprefinterval {
		msgsplit := strings.Split(msg, " ")
		answermas = service.GetWeatherByNDays(msgsplit[1], chatId)
		return true, answermas
	}
	return false, answermas
}

func Calculator(msg string, chatId int64, person map[int64]rune) bool {
	checkprefinterval, _ := regexp.MatchString("calc [-+*/]", msg)
	if checkprefinterval {
		symbol := msg[len(msg)-1]
		person[chatId] = rune(symbol)
		return true
	}
	return false
}

func isTwoNumbers(msg string, chatId int64, person map[int64]rune) string {
	var rez float64
	var answer string
	checkprefinterval, _ := regexp.MatchString("[-]?\\d+ [-]?\\d+", msg)
	if checkprefinterval {
		msgsplit := strings.Split(msg, " ")
		num1, _ := strconv.Atoi(msgsplit[0])
		num2, _ := strconv.Atoi(msgsplit[1])
		switch person[chatId] {
		case '+':
			rez = float64(num1) + float64(num2)
		case '-':
			rez = float64(num1) - float64(num2)
		case '*':
			rez = float64(num1) * float64(num2)
		case '/':
			rez = float64(num1) / float64(num2)
		default:
			return "Не выбрали операцию в калькуляторе"
		}
		delete(person, chatId)
		answer = strconv.FormatFloat(rez, 'f', 1, 64)
		return answer
	}
	return answer
}

func CalFruit(msg string) (bool, string) {
	var answer string
	checkprefdata, _ := regexp.MatchString("apiFruit .+", msg)
	if checkprefdata {
		msgsplit := strings.Split(msg, " ")
		fruitinfo := parse.ParseFruit(msgsplit[1])
		answer = "Каллорийность: " + strconv.FormatFloat(fruitinfo.Calories, 'f', 1, 64) +
			"\nЖиры: " + strconv.FormatFloat(fruitinfo.Fats, 'f', 1, 64) +
			"\nСахар: " + strconv.FormatFloat(fruitinfo.Sugar, 'f', 1, 64) +
			"\nУглеводы: " + strconv.FormatFloat(fruitinfo.Carbohyddrates, 'f', 1, 64) +
			"\nБелок: " + strconv.FormatFloat(fruitinfo.Protein, 'f', 1, 64)
		return true, answer
	}
	return false, answer
}

func Btc(msg string) (bool, string) {
	var answer string
	checkprefBtc, _ := regexp.MatchString("BTC/USD", msg)
	if checkprefBtc {
		answer = parse.ParseBtc()
		return true, answer
	}
	return false, answer
}
