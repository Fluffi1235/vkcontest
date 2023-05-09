package bot

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
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

func (b *Bot) RunBot(ctx context.Context, wg *sync.WaitGroup) {
	msgChan := make(chan *model.Message)

	for _, source := range b.Sources {
		go source.Read(ctx, msgChan)
	}

	b.HandlingMessage(msgChan)

	close(msgChan)
	wg.Done()
}

func (b *Bot) HandlingMessage(msgChan <-chan *model.Message) {
	mpperson := make(map[int64]rune, 0)
	var answer string
	for msg := range msgChan {
		if msg.Button != nil {
			log.Printf("Нажата кнопка с данными: %s\n", msg.Button.Data)

			if funcResponse, answerdata := DataUser(msg.Button.Data, msg.ChatID); funcResponse {
				b.Sources[msg.Source].EditMessage(answerdata, msg.ChatID, msg.Button.Message.MessageID)
			}

			if funcResponse, answercity := CheckCity(msg.Button.Data, msg.ChatID); funcResponse {
				b.Sources[msg.Source].EditMessage(answercity, msg.ChatID, msg.Button.Message.MessageID)
			}

			if funcResponse, answermasNday := CheckNdays(msg.Button.Data, msg.ChatID); funcResponse {
				for i := 1; i < len(answermasNday); i++ {
					answer = answer + answermasNday[i]
					if i%4 == 0 && i < 5 {
						b.Sources[msg.Source].EditMessage(answer, msg.ChatID, msg.Button.Message.MessageID)
						answer = ""
					}
					if i%4 == 0 && i >= 5 {
						b.Sources[msg.Source].Send(answer, msg.ChatID)
						answer = ""
					}
				}
			}

			if Calculator(msg.Button.Data, msg.ChatID, mpperson) {
				b.Sources[msg.Source].EditMessage("Введите 2 числа через пробел", msg.ChatID, msg.Button.Message.MessageID)
			}

			if funcResponse, answerCalFruit := CalFruit(msg.Button.Data); funcResponse {
				b.Sources[msg.Source].EditMessage(answerCalFruit, msg.ChatID, msg.Button.Message.MessageID)
			}

			if funcResponse, answerBTC := Btc(msg.Button.Data); funcResponse {
				b.Sources[msg.Source].EditMessage("Текущий курс BTC/USD: "+answerBTC+"\nДанные были взяты с сайта https://www.coindesk.com/search?s=bitcoin", msg.ChatID, msg.Button.Message.MessageID)
			}

			switch msg.Button.Data {
			case "Мои данные":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Изменить город":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Прогноз погоды":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Калькулятор":
				b.Sources[msg.Source].EditMessageWithButtons("Мои функции выберите одну из них", msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "OPEN API":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Калорийность фруктов":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
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
				service.RegistrUser(msg)
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
