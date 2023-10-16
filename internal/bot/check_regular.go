package bot

import (
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/services"
)

func CheckRegular(b *Bot, msg *model.MessageInfoButton, service *services.Repository, persons map[int64]rune, config *config.Config) bool {
	if answer, err := DataUser(msg.ButtonData, msg.ButtonInfo.ChatID, service, msg.ButtonInfo.Platform); answer != "" && err == nil {
		editMessage := model.NewEditMessage(answer, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if answer, err := CheckCity(msg.ButtonData, msg.ButtonInfo.ChatID, service); answer != "" && err == nil {
		editMessage := model.NewEditMessage(answer, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if arrAnswer, err := WeatherOnNDays(msg.ButtonData, msg.ButtonInfo.ChatID, service); len(arrAnswer) > 0 && err == nil {
		for i := 1; i < len(arrAnswer); i++ {
			if i == 1 {
				editMessage := model.NewEditMessage(arrAnswer[i], msg.ButtonInfo.ChatID, msg.ButtonMessageID)
				b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
			} else if i > 1 {
				b.Sources[msg.ButtonInfo.Source].Send(arrAnswer[i], msg.ButtonInfo.ChatID)
			}
		}
		return true
	}
	if answerCalc, err := Calculator(msg.ButtonData, msg.ButtonInfo.ChatID, persons); answerCalc && err == nil {
		editMessage := model.NewEditMessage("Введите 2 числа через пробел", msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if answerCalFruit, err := InfoFruits(msg.ButtonData, config); answerCalFruit != "" && err == nil {
		editMessage := model.NewEditMessage(answerCalFruit, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if answerBTC, err := Btc(msg.ButtonData, config); answerBTC != "" && err == nil {
		editMessage := model.NewEditMessage(
			fmt.Sprintf("%s%s%s", "Текущий курс BTC/USD: ", answerBTC, "\n\nДанные были взяты с сайта https://www.coindesk.com/search?s=bitcoin"),
			msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	return false
}
