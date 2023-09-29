package bot

import (
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/services"
	"log"
	"strings"
)

func (b *Bot) HandlingButton(msg *model.MessageInfoButton, service *services.Repository, persons map[int64]rune) {
	if msg.ButtonData != "" && msg.ButtonMessageID != -1 {
		log.Printf("Нажата кнопка с данными: %s\n", msg.ButtonData)
		msg.ButtonData = strings.ToLower(msg.ButtonData)
		if funcResponse, answerdata := services.DataUser(msg.ButtonData, msg.Mi.ChatID, service, msg.Mi.Platform); funcResponse {
			editmessage := model.NewEditMessage(answerdata, msg.Mi.ChatID, msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessage(editmessage)
		}
		if funcResponse, answercity := services.CheckCity(msg.ButtonData, msg.Mi.ChatID, service); funcResponse {
			editmessage := model.NewEditMessage(answercity, msg.Mi.ChatID, msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessage(editmessage)
		}
		if funcResponse, answermasNday := services.CheckNdays(msg.ButtonData, msg.Mi.ChatID, service); funcResponse {
			var answer string
			for i := 1; i < len(answermasNday); i++ {
				answer = answer + answermasNday[i]
				if i%4 == 0 && i < 5 {
					editmessage := model.NewEditMessage(answer, msg.Mi.ChatID, msg.ButtonMessageID)
					b.Sources[msg.Mi.Source].EditMessage(editmessage)
					answer = ""
				}
				if i%4 == 0 && i >= 5 {
					b.Sources[msg.Mi.Source].Send(answer, msg.Mi.ChatID)
					answer = ""
				}
			}
		}
		if services.Calculator(msg.ButtonData, msg.Mi.ChatID, persons) {
			editmessage := model.NewEditMessage("Введите 2 числа через пробел", msg.Mi.ChatID, msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessage(editmessage)
		}
		if funcResponse, answerCalFruit := services.CalFruit(msg.ButtonData); funcResponse {
			editmessage := model.NewEditMessage(answerCalFruit, msg.Mi.ChatID, msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessage(editmessage)
		}
		if funcResponse, answerBTC := services.Btc(msg.ButtonData); funcResponse {
			editmessage := model.NewEditMessage(
				fmt.Sprintf("%s%s%s", "Текущий курс BTC/USD: ", answerBTC, "\nДанные были взяты с сайта https://www.coindesk.com/search?s=bitcoin"),
				msg.Mi.ChatID, msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessage(editmessage)
		}
		switch msg.ButtonData {
		case "мои данные":
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.Mi.ChatID,
				msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessageWithButtons(editMessageWithButtons)
		case "изменить город":
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.Mi.ChatID,
				msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessageWithButtons(editMessageWithButtons)
		case "прогноз погоды":
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.Mi.ChatID,
				msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessageWithButtons(editMessageWithButtons)
		case "калькулятор":
			editMessageWithButtons := model.NewEditMessageWithButtons("Мои функции выберите одну из них", msg.ButtonData, msg.Mi.ChatID, msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessageWithButtons(editMessageWithButtons)
		case "open api":
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.Mi.ChatID,
				msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessageWithButtons(editMessageWithButtons)
		case "калорийность фруктов":
			editMessageWithButtons := model.NewEditMessageWithButtons(fmt.Sprintf("%s%s", "Вы нажали кнопку ", msg.ButtonData), msg.ButtonData, msg.Mi.ChatID,
				msg.ButtonMessageID)
			b.Sources[msg.Mi.Source].EditMessageWithButtons(editMessageWithButtons)
		}
	}
}

func (b *Bot) HandlingText(msg *model.MessageInfoText, service *services.Repository, persons map[int64]rune, repo repository.UniversalRepo) {
	if msg.Text != "" {
		message := strings.ToLower(msg.Text)
		answer := services.IsTwoNumbers(message, msg.Mi.ChatID, persons)
		if answer != "" {
			b.Sources[msg.Mi.Source].Send(answer, msg.Mi.ChatID)
			return
		}
		switch message {
		case "/start":
			repo.RegistrUser(msg)
			answer = service.AnswerStart()
			b.Sources[msg.Mi.Source].Send(answer, msg.Mi.ChatID)
		case "начать":
			repo.RegistrUser(msg)
			answer = service.AnswerStart()
			b.Sources[msg.Mi.Source].Send(answer, msg.Mi.ChatID)
		case "/info":
			answer = "Выберите функцию"
			b.Sources[msg.Mi.Source].SendButton(answer, msg.Mi.ChatID)
		case "/help":
			answer = "Чтобы узнать какие команды есть введите /info"
			b.Sources[msg.Mi.Source].Send(answer, msg.Mi.ChatID)
		default:
			b.Sources[msg.Mi.Source].Send("Что то пошло не так, попробуйте еще раз /info", msg.Mi.ChatID)
		}
	}
}
