package bot

import (
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/services"
	"log"
	"strings"
)

const (
	userData        string = "мои данные"
	changeCity      string = "изменить город"
	weatherForecast string = "прогноз погоды"
	calculator      string = "калькулятор"
	openAPI         string = "open api"
	fruitInfo       string = "калорийность фруктов"
)

func (b *Bot) HandlingButton(msg *model.MessageInfoButton, service *services.Repository, persons map[int64]rune, config *config.Config) error {
	if msg.ButtonData != "" && msg.ButtonMessageID != -1 {
		log.Printf("[%s]Нажата кнопка с данными: %s\n", msg.ButtonInfo.Source, msg.ButtonData)
		msg.ButtonData = strings.ToLower(msg.ButtonData)
		err := b.CheckRegular(msg, service, persons, config)
		if err != nil {
			return err
		}
		switch msg.ButtonData {
		case userData:
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.ButtonInfo.ChatID,
				msg.ButtonMessageID)
			err = b.Sources[msg.ButtonInfo.Source].EditMessageWithButtons(editMessageWithButtons)
			if err != nil {
				return err
			}
		case changeCity:
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.ButtonInfo.ChatID,
				msg.ButtonMessageID)
			err = b.Sources[msg.ButtonInfo.Source].EditMessageWithButtons(editMessageWithButtons)
			if err != nil {
				return err
			}
		case weatherForecast:
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.ButtonInfo.ChatID,
				msg.ButtonMessageID)
			err = b.Sources[msg.ButtonInfo.Source].EditMessageWithButtons(editMessageWithButtons)
			if err != nil {
				return err
			}
		case calculator:
			editMessageWithButtons := model.NewEditMessageWithButtons("Мои функции, выберите одну из них", msg.ButtonData, msg.ButtonInfo.ChatID, msg.ButtonMessageID)
			err = b.Sources[msg.ButtonInfo.Source].EditMessageWithButtons(editMessageWithButtons)
			if err != nil {
				return err
			}
		case openAPI:
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.ButtonInfo.ChatID,
				msg.ButtonMessageID)
			err = b.Sources[msg.ButtonInfo.Source].EditMessageWithButtons(editMessageWithButtons)
			if err != nil {
				return err
			}
		case fruitInfo:
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.ButtonInfo.ChatID,
				msg.ButtonMessageID)
			err = b.Sources[msg.ButtonInfo.Source].EditMessageWithButtons(editMessageWithButtons)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
