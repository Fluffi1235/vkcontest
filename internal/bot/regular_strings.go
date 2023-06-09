package bot

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parse"
	"github.com/Fluffi1235/vkcontest/internal/service"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func DataUser(msg string, chatId int64, servise *service.Repository, platform string) (bool, string) {
	var answer string

	checkprefdata, _ := regexp.MatchString("показать мои данные", msg)
	if checkprefdata {
		answer = servise.GetUserData(chatId, platform)
		return true, answer
	}
	return false, answer
}

func CheckCity(msg string, chatId int64, servise *service.Repository) (bool, string) {
	citys := model.City()
	var city string
	var answer string
	for k, _ := range citys {
		checkprefday, err := regexp.MatchString("city .+", msg)
		if err != nil {
			log.Println(err)
		}
		if checkprefday {
			city = strings.Split(msg, "city ")[1]
			if city == k {
				answer = servise.AnswerForCityChange(city)
				servise.SaveCity(k, chatId)
				return true, answer
			} else {
				answer = "Информации о погоде об этом городе нет"
			}
		}
	}
	return false, answer
}

func CheckNdays(msg string, chatId int64, servise *service.Repository) (bool, []string) {
	answermas := make([]string, 0)
	checkprefinterval, err := regexp.MatchString("погода \\d{1,10}", msg)
	if err != nil {
		log.Println(err)
	}
	if checkprefinterval {
		msgsplit := strings.Split(msg, " ")
		answermas = servise.GetWeatherByNDays(msgsplit[1], chatId)
		return true, answermas
	}
	return false, answermas
}

func Calculator(msg string, chatId int64, person map[int64]rune) bool {
	checkprefinterval, err := regexp.MatchString("calc [-+*/]", msg)
	if err != nil {
		log.Println(err)
	}
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
	checkprefinterval, err := regexp.MatchString("[-]?\\d+[.,]?\\d* [-]?\\d+[.,]?\\d*", msg)
	if err != nil {
		log.Println(err)
	}
	if checkprefinterval {
		msgsplit := strings.Split(msg, " ")
		msgsplit[0] = strings.ReplaceAll(msgsplit[0], ",", ".")
		msgsplit[1] = strings.ReplaceAll(msgsplit[1], ",", ".")
		num1, _ := strconv.ParseFloat(msgsplit[0], 64)
		num2, _ := strconv.ParseFloat(msgsplit[1], 64)
		switch person[chatId] {
		case '+':
			rez = num1 + num2
		case '-':
			rez = num1 - num2
		case '*':
			rez = num1 * num2
		case '/':
			if num2 == 0 {
				return "На 0 делить нельзя"
			}
			rez = num1 / num2
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
	checkprefdata, err := regexp.MatchString("apifruit .+", msg)
	if err != nil {
		log.Println(err)
	}
	if checkprefdata {
		msgsplit := strings.Split(msg, " ")
		fruitinfo := parse.ParseFruit(msgsplit[1])
		answer = "Энергетическая ценность " + msgsplit[2] + " на 100г: " +
			"\nКалорийность: " + strconv.FormatFloat(fruitinfo.Calories, 'f', 1, 64) + "кк" +
			"\nПищевая ценность " + msgsplit[2] + " на 100г:" +
			"\nЖиры: " + strconv.FormatFloat(fruitinfo.Fats, 'f', 1, 64) + "г" +
			"\nСахар: " + strconv.FormatFloat(fruitinfo.Sugar, 'f', 1, 64) + "г" +
			"\nУглеводы: " + strconv.FormatFloat(fruitinfo.Carbohyddrates, 'f', 1, 64) + "г" +
			"\nБелки: " + strconv.FormatFloat(fruitinfo.Protein, 'f', 1, 64) + "г"
		return true, answer
	}
	return false, answer
}

func Btc(msg string) (bool, string) {
	var answer string
	checkprefBtc, err := regexp.MatchString("btc", msg)
	if err != nil {
		log.Println(err)
	}
	if checkprefBtc {
		answer = parse.ParseBtc()
		return true, answer
	}
	return false, answer
}
