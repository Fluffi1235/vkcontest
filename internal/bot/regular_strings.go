package bot

import (
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parsers"
	"github.com/Fluffi1235/vkcontest/internal/services"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

func DataUser(msg string, chatId int64, servise *services.Repository, platform string) (string, error) {
	var answer string

	checkPrefData, err := regexp.MatchString("показать мои данные", msg)
	if err != nil {
		return "", errors.Wrap(err, "[Regular DataUser]")

	}
	if checkPrefData {
		answer, err = servise.AnswerUserData(chatId, platform)
		if err != nil {
			return "", err
		}
		return answer, nil
	}
	return answer, nil
}

func CheckCity(msg string, chatId int64, servise *services.Repository) (string, error) {
	cities := model.Cities()
	var city string
	var answer string
	for k := range cities {
		checkPrefDay, err := regexp.MatchString("city .+", msg)
		if err != nil {
			return "", err
		}
		if checkPrefDay {
			city = strings.Split(msg, "city ")[1]
			if city == k {
				answer, err = servise.AnswerForCityChange(city, chatId)
				if err != nil {
					return "", err
				}
				return answer, nil
			} else {
				answer = "Информации о погоде об этом городе нет"
			}
		}
	}
	return answer, nil
}

func WeatherOnNDays(msg string, chatId int64, servise *services.Repository) ([]string, error) {
	arrAnswer := make([]string, 0)
	checkPrefInterval, err := regexp.MatchString("погода \\d{1,10}", msg)
	if err != nil {
		return nil, err
	}
	if checkPrefInterval {
		msgSplit := strings.Split(msg, " ")
		arrAnswer, err = servise.AnswerWeatherByNDays(msgSplit[1], chatId)
		if err != nil {
			return nil, err
		}
	}
	return arrAnswer, nil
}

func Calculator(msg string, chatId int64, person map[int64]rune) (bool, error) {
	checkPrefCalc, err := regexp.MatchString("calc [-+*/]", msg)
	if err != nil {
		return false, err
	}
	if checkPrefCalc {
		symbol := msg[len(msg)-1]
		person[chatId] = rune(symbol)
		return true, nil
	}
	return false, nil
}

func IsTwoNumbers(msg string, chatId int64, person map[int64]rune) (string, error) {
	var res float64
	var answer string
	var num1 float64
	var num2 float64
	checkPrefTwoNumbers, err := regexp.MatchString("[-]?\\d+[.,]?\\d* [-]?\\d+[.,]?\\d*", msg)
	if err != nil {
		return "", err
	}
	if checkPrefTwoNumbers {
		msgSplit := strings.Split(msg, " ")
		msgSplit[0] = strings.ReplaceAll(msgSplit[0], ",", ".")
		msgSplit[1] = strings.ReplaceAll(msgSplit[1], ",", ".")
		num1, err = strconv.ParseFloat(msgSplit[0], 64)
		if err != nil {
			return "", err
		}
		num2, err = strconv.ParseFloat(msgSplit[1], 64)
		if err != nil {
			return "", err
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
				return "На 0 делить нельзя", nil
			}
			res = num1 / num2
		default:
			return "Не выбрали операцию в калькуляторе", nil
		}
		delete(person, chatId)
		answer = strconv.FormatFloat(res, 'f', 1, 64)
	}
	return answer, nil
}

func InfoFruits(msg string, config *config.Config) (string, error) {
	var answer string
	checkPrefFruit, err := regexp.MatchString("apifruit .+", msg)
	if err != nil {
		return "", err
	}
	if checkPrefFruit {
		fruit := strings.Split(msg, " ")
		fruitInfo, err := parsers.ParseFruit(fruit[1], config)
		if err != nil {
			return "", err
		}
		answer, err = services.AnswerInfoFruits(fruit[2], fruitInfo)
		if err != nil {
			return "", err
		}
	}
	return answer, nil
}

func Btc(msg string, config *config.Config) (string, error) {
	var answer string
	checkPrefBtc, err := regexp.MatchString("btc", msg)
	if err != nil {
		return "", err
	}
	if checkPrefBtc {
		answer, err = parsers.ParseBtc(config)
		if err != nil {
			return "", err
		}
	}
	return answer, nil
}
