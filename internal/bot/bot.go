package bot

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/service"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	"log"
	"strings"
	"sync"
)

type Bot struct {
	Sources map[model.SourceType]sources.Source
}

func NewBot(m map[model.SourceType]sources.Source) Bot {
	return Bot{
		Sources: m,
	}
}

func (b *Bot) RunBot(ctx context.Context, wg *sync.WaitGroup, repo repository.UniversalRepo) {
	msgChan := make(chan *model.Message)

	for _, source := range b.Sources {
		go source.Read(ctx, msgChan)
	}

	b.HandlingMessage(msgChan, repo)

	close(msgChan)
	wg.Done()
}

func (b *Bot) HandlingMessage(msgChan <-chan *model.Message, repo repository.UniversalRepo) {
	service := service.New(repo)
	mpperson := make(map[int64]rune, 0)
	var answer string
	for msg := range msgChan {
		if msg.ButtonDate != "" && msg.ButtonMessageID != 0 {
			log.Printf("Нажата кнопка с данными: %s\n", msg.ButtonDate)

			if funcResponse, answerdata := DataUser(msg.ButtonDate, msg.ChatID, service, msg.Platform); funcResponse {
				b.Sources[msg.Source].EditMessage(answerdata, msg.ChatID, msg.ButtonMessageID)
			}

			if funcResponse, answercity := CheckCity(msg.ButtonDate, msg.ChatID, service); funcResponse {
				b.Sources[msg.Source].EditMessage(answercity, msg.ChatID, msg.ButtonMessageID)
			}

			if funcResponse, answermasNday := CheckNdays(msg.ButtonDate, msg.ChatID, service); funcResponse {
				for i := 1; i < len(answermasNday); i++ {
					answer = answer + answermasNday[i]
					if i%4 == 0 && i < 5 {
						b.Sources[msg.Source].EditMessage(answer, msg.ChatID, msg.ButtonMessageID)
						answer = ""
					}
					if i%4 == 0 && i >= 5 {
						b.Sources[msg.Source].Send(answer, msg.ChatID)
						answer = ""
					}
				}
			}

			if Calculator(msg.ButtonDate, msg.ChatID, mpperson) {
				b.Sources[msg.Source].EditMessage("Введите 2 числа через пробел", msg.ChatID, msg.ButtonMessageID)
			}

			if funcResponse, answerCalFruit := CalFruit(msg.ButtonDate); funcResponse {
				b.Sources[msg.Source].EditMessage(answerCalFruit, msg.ChatID, msg.ButtonMessageID)
			}

			if funcResponse, answerBTC := Btc(msg.ButtonDate); funcResponse {
				b.Sources[msg.Source].EditMessage("Текущий курс BTC/USD: "+answerBTC+
					"\nДанные были взяты с сайта https://www.coindesk.com/search?s=bitcoin", msg.ChatID, msg.ButtonMessageID)
			}

			switch msg.ButtonDate {
			case "Мои данные":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.ButtonDate, msg.ChatID, msg.ButtonDate, msg.ButtonMessageID)
			case "Изменить город":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.ButtonDate, msg.ChatID, msg.ButtonDate, msg.ButtonMessageID)
			case "Прогноз погоды":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.ButtonDate, msg.ChatID, msg.ButtonDate, msg.ButtonMessageID)
			case "Калькулятор":
				b.Sources[msg.Source].EditMessageWithButtons("Мои функции выберите одну из них", msg.ChatID, msg.ButtonDate, msg.ButtonMessageID)
			case "OPEN API":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.ButtonDate, msg.ChatID, msg.ButtonDate, msg.ButtonMessageID)
			case "Калорийность фруктов":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.ButtonDate, msg.ChatID, msg.ButtonDate, msg.ButtonMessageID)
			}
		}
		if msg.Text != "" {
			message := strings.ToLower(msg.Text)

			answer = isTwoNumbers(message, msg.ChatID, mpperson)
			if answer != "" {
				b.Sources[msg.Source].Send(answer, msg.ChatID)
			}

			switch message {
			case "/start":
				repo.RegistrUser(msg)
				answer = service.AnswerStart()
				b.Sources[msg.Source].Send(answer, msg.ChatID)
			case "/info":
				answer = "Выберите функцию"
				b.Sources[msg.Source].SendButton(answer, msg.ChatID)
			case "/help":
				answer = "Чтобы узнать какие команды есть введите /info"
				b.Sources[msg.Source].Send(answer, msg.ChatID)
			}
			if answer == "" {
				b.Sources[msg.Source].Send("Что то пошло не так, попробуйте еще раз /info", msg.ChatID)
			}
		}
	}
}
