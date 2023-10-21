package bot

import (
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/services"
)

func (b *Bot) CheckRegular(msg *model.MessageInfoButton, service *services.Repository, persons map[int64]rune, config *config.Config) (bool, error) {
	if answer, err := DataUser(msg.ButtonData, msg.ButtonInfo.ChatID, service, msg.ButtonInfo.Platform); answer != "" {
		editMessage := model.NewEditMessage(answer, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		err = b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true, err
	} else if err != nil {
		return true, err
	}
	if answer, err := CheckCity(msg.ButtonData, msg.ButtonInfo.ChatID, service); answer != "" {
		editMessage := model.NewEditMessage(answer, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		err = b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true, err
	} else if err != nil {
		return true, err
	}
	if arrAnswer, err := WeatherOnNDays(msg.ButtonData, msg.ButtonInfo.ChatID, service); len(arrAnswer) > 0 {
		for i := 1; i < len(arrAnswer); i++ {
			if i == 1 {
				editMessage := model.NewEditMessage(arrAnswer[i], msg.ButtonInfo.ChatID, msg.ButtonMessageID)
				err = b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
			} else if i > 1 {
				err = b.Sources[msg.ButtonInfo.Source].Send(arrAnswer[i], msg.ButtonInfo.ChatID)
			}
			if err != nil {
				break
			}
		}
		return true, err
	} else if err != nil {
		return true, err
	}
	if answerCalc, err := Calculator(msg.ButtonData, msg.ButtonInfo.ChatID, persons); answerCalc && err != nil {
		editMessage := model.NewEditMessage("Введите 2 числа через пробел\nПример: 3 5", msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		err = b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true, err
	} else if err != nil {
		return true, err
	}
	if answerCalFruit, err := InfoFruits(msg.ButtonData, config); answerCalFruit != "" {
		editMessage := model.NewEditMessage(answerCalFruit, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		err = b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true, err
	} else if err != nil {
		return true, err
	}
	if answerBTC, err := Btc(msg.ButtonData, config); answerBTC != "" {
		editMessage := model.NewEditMessage(
			fmt.Sprintf("%s%s%s", "Текущий курс BTC/USD: ", answerBTC, "\n\nДанные были взяты с сайта https://www.coindesk.com/search?s=bitcoin"),
			msg.ButtonInfo.ChatID, msg.ButtonMessageID)
		err = b.Sources[msg.ButtonInfo.Source].EditMessage(editMessage)
		return true, err
	} else if err != nil {
		return true, err
	}
	return false, nil
}
