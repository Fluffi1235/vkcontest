package bot

import (
	"context"
	"log"
	"strings"
	"sync"
	"telegram_bot/internal/model"
	"telegram_bot/internal/service"
	"telegram_bot/internal/sources"
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
			textButton := msg.Button.Data

			if Calculator(msg.Button.Data, msg.ChatID, mpperson) {
				b.Sources[msg.Source].EditMessage("Введите 2 числа через пробел", msg.ChatID, msg.Button.Message.MessageID)
			}

			if a, answercity := CheckCity(msg.Button.Data, msg.ChatID); a == true {
				b.Sources[msg.Source].EditMessage(answercity, msg.ChatID, msg.Button.Message.MessageID)
			}

			if a, answermasNday := CheckNdays(msg.Button.Data, msg.ChatID); a == true {
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
			switch textButton {
			case "Мои данные":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Прогноз погоды":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Калькулятор":
				b.Sources[msg.Source].EditMessageWithButtons("Мои функции выберите одну из них", msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "OPEN API":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Калорийность фруктов":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			case "Изменить город":
				b.Sources[msg.Source].EditMessageWithButtons("Вы нажали кнопку "+msg.Button.Data, msg.ChatID, msg.Button.Data, msg.Button.Message.MessageID)
			}
		}
		if msg.Text != "" {
			message := strings.ToLower(msg.Text)

			answer = isTwoNumbers(msg.Text, msg.ChatID, mpperson)
			if answer != "" {
				b.Sources[msg.Source].Send(answer, msg.ChatID)
			}

			switch message {

			case "/start":
				service.RegistrUser(msg)
				b.Sources[msg.Source].Send(service.AnswerStart(), msg.ChatID)

			case "/info":
				b.Sources[msg.Source].SendButton("Выберите функцию", msg.ChatID)

			case "/help":
				answer = "Чтобы узнать какие команды есть введите /info"
				b.Sources[msg.Source].Send(answer, msg.ChatID)
			}
		}
	}
}
