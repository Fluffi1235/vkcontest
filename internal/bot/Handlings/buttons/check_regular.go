package buttons

import (
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/bot"
	"github.com/Fluffi1235/vkcontest/internal/load_configs"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/services"
)

func CheckRegular(b *bot.Bot, msg *model.MessageInfoButton, service *services.Repository, persons map[int64]rune, config *load_configs.Config) bool {
	if funcResponse, answer := DataUser(msg.ButtonData, msg.ButtonInfo.ChatID, service, msg.ButtonInfo.Platform); funcResponse {
		editMessage := model.NewEditMessage(answer, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if funcResponse, answer := CheckCity(msg.ButtonData, msg.ButtonInfo.ChatID, service); funcResponse {
		editMessage := model.NewEditMessage(answer, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if funcResponse, arrAnswer := WeatherOnNDays(msg.ButtonData, msg.ButtonInfo.ChatID, service); funcResponse {
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
	if Calculator(msg.ButtonData, msg.ButtonInfo.ChatID, persons) {
		editMessage := model.NewEditMessage("Введите 2 числа через пробел", msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if funcResponse, answerCalFruit := InfoFruits(msg.ButtonData, config); funcResponse {
		editMessage := model.NewEditMessage(answerCalFruit, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	if funcResponse, answerBTC := Btc(msg.ButtonData, config); funcResponse {
		editMessage := model.NewEditMessage(
			fmt.Sprintf("%s%s%s", "Текущий курс BTC/USD: ", answerBTC, "\n\nДанные были взяты с сайта https://www.coindesk.com/search?s=bitcoin"),
			msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true
	}
	return false
}
