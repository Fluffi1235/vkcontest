package bot

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"telegram_bot/internal/model"
	"telegram_bot/internal/service"
)

func CheckCity(msg string, chatid int64) (bool, string) {
	citys := model.City()
	city := ""
	answer := ""
	for k, _ := range citys {
		checkprefday, _ := regexp.MatchString("city .+", msg)
		if checkprefday {
			city = strings.Split(msg, " ")[1]
			if city == k {
				answer = service.AnswerForCheakprefCity(city)
				service.SaveCity(k, chatid)
				return true, answer
			} else {
				answer = "Информации о погоде об этом городе нет"
			}
		}
	}
	return false, answer
}

func Checkday(msg string, chatid int64) string {
	answer := ""
	checkprefday, _ := regexp.MatchString("погода на \\d{4}-\\d{2}-\\d{2}", msg)
	if checkprefday {
		msgsplit := strings.Split(msg, " ")
		answer = service.GetWeatherByDate(msgsplit[2], chatid)
		if answer == "" {
			answer = "Невалидная дата"
		}
	}
	return answer
}

func CheckInterval(msg string, chatid int64) ([]string, error) {
	answermas := make([]string, 0)
	checkprefinterval, _ := regexp.MatchString("погода с \\d{4}-\\d{2}-\\d{2} по \\d{4}-\\d{2}-\\d{2}", msg)
	if checkprefinterval {
		msgsplit := strings.Split(msg, " ")
		answermas = service.GetWeatherByInterval(msgsplit[2], msgsplit[4], chatid)
	}
	if len(answermas) == 1 {
		return answermas, errors.New("Таких дат нет")
	}
	return answermas, nil
}

func CheckNdays(msg string, chatid int64) (bool, []string) {
	answermas := make([]string, 0)
	checkprefinterval, _ := regexp.MatchString("Погода \\d{1,10}", msg)
	if checkprefinterval {
		msgsplit := strings.Split(msg, " ")
		answermas = service.GetWeatherByNDays(msgsplit[1], chatid)
		return true, answermas
	}
	return false, answermas
}

func Calculator(msg string, chatid int64, person map[int64]rune) bool {
	checkprefinterval, _ := regexp.MatchString("calc [-+*/]", msg)
	if checkprefinterval {
		symbol := msg[len(msg)-1]
		person[chatid] = rune(symbol)
		return true
	}
	return false
}

func isTwoNumbers(msg string, chatid int64, person map[int64]rune) string {
	var rez float64
	var answer string
	checkprefinterval, _ := regexp.MatchString("\\d+ \\d+", msg)
	if checkprefinterval {
		msgsplit := strings.Split(msg, " ")
		num1, _ := strconv.Atoi(msgsplit[0])
		num2, _ := strconv.Atoi(msgsplit[1])
		switch person[chatid] {
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
		delete(person, chatid)
		answer = strconv.FormatFloat(rez, 'f', -1, 64)
		return answer
	}
	return ""
}
