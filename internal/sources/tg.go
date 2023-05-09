package sources

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Source interface {
	Read(ctx context.Context, msgChan chan<- *model.Message)
	Send(msg string, clientID int64)
	GetSource() model.SourceType
	SendButton(msg string, clientID int64)
	EditMessageWithButtons(msg string, clientID int64, button string, msgId int)
	EditMessage(msg string, clientID int64, msgId int)
}

type TG struct {
	Chan   tgbotapi.UpdatesChannel
	bot    *tgbotapi.BotAPI
	CharID int64
}

func NewTG(token string) Source {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	tg := &TG{
		Chan: updates,
		bot:  bot,
	}

	return tg
}

func (tg *TG) Read(ctx context.Context, msgChan chan<- *model.Message) {
	for update := range tg.Chan {
		print()

		select {
		case <-ctx.Done():
			return
		default:
		}

		if update.Message != nil { // If we got a message
			msg := &model.Message{
				Source:    tg.GetSource(),
				Text:      update.Message.Text,
				ChatID:    update.Message.Chat.ID,
				Username:  update.Message.Chat.UserName,
				FirstName: update.Message.Chat.FirstName,
				LastName:  update.Message.Chat.LastName,
			}
			_ = msg

			msgChan <- msg
		}
		if update.CallbackQuery != nil {
			msg := &model.Message{
				Source: tg.GetSource(),
				ChatID: int64(update.CallbackQuery.From.ID),
				Button: update.CallbackQuery,
			}
			_ = msg

			msgChan <- msg
		}
	}
}

func (tg *TG) GetSource() model.SourceType {
	return model.Telegram
}

func (tg *TG) Send(msg string, clientID int64) {
	tgMsg := tgbotapi.NewMessage(clientID, msg)

	_, err := tg.bot.Send(tgMsg)
	if err != nil {
		log.Printf("%v", err)
	}
}

func (tg *TG) SendButton(msg string, clientID int64) {
	tgMsg := tgbotapi.NewMessage(clientID, msg)
	tgMsg.ReplyMarkup = createInlineKeyboardInfo()

	_, err := tg.bot.Send(tgMsg)
	if err != nil {
		log.Fatal(err)
	}

}

func (tg *TG) EditMessageWithButtons(msg string, clientID int64, button string, msgId int) {
	editmsg := tgbotapi.NewEditMessageText(clientID, msgId, msg)
	switch button {
	case "Мои данные":
		editmsg.ReplyMarkup = createInlineKeyboardData()
	case "Прогноз погоды":
		editmsg.ReplyMarkup = createInlineKeyboardWeather()
	case "Калькулятор":
		editmsg.ReplyMarkup = createInlineKeyboardCalculator()
	case "OPEN API":
		editmsg.ReplyMarkup = createInlineKeyboardOpenAPI()
	case "Калорийность фруктов":
		editmsg.ReplyMarkup = createInlineKeyboardFruits()
	case "Изменить город":
		editmsg.ReplyMarkup = createInlineKeyboardCity()
	}
	_, err := tg.bot.Send(editmsg)
	if err != nil {
		log.Println(err)
	}
}

func (tg *TG) EditMessage(msg string, clientID int64, msgId int) {
	editmsg := tgbotapi.NewEditMessageText(clientID, msgId, msg)
	_, err := tg.bot.Send(editmsg)
	if err != nil {
		log.Println(err)
	}
}

func createInlineKeyboardInfo() tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Мои данные", "Мои данные")
	keyboardButtons = append(keyboardButtons, button1)
	button2 := tgbotapi.NewInlineKeyboardButtonData("Прогноз погоды", "Прогноз погоды")
	keyboardButtons = append(keyboardButtons, button2)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	button3 := tgbotapi.NewInlineKeyboardButtonData("Калькулятор", "Калькулятор")
	keyboardButtons = append(keyboardButtons, button3)
	button4 := tgbotapi.NewInlineKeyboardButtonData("OPEN API", "OPEN API")
	keyboardButtons = append(keyboardButtons, button4)

	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[2:]...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	return keyboard
}

func createInlineKeyboardData() *tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Показать мои данные", "Показать мои данные")
	keyboardButtons = append(keyboardButtons, button1)
	button2 := tgbotapi.NewInlineKeyboardButtonData("Изменить город", "Изменить город")
	keyboardButtons = append(keyboardButtons, button2)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1)

	return &keyboard
}

func createInlineKeyboardCity() *tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Санкт-Петербург", "city санкт-петербург")
	keyboardButtons = append(keyboardButtons, button1)
	button2 := tgbotapi.NewInlineKeyboardButtonData("Москва", "city москва")
	keyboardButtons = append(keyboardButtons, button2)
	button3 := tgbotapi.NewInlineKeyboardButtonData("Новосибирск", "city новосибирск")
	keyboardButtons = append(keyboardButtons, button3)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	button4 := tgbotapi.NewInlineKeyboardButtonData("Екатеренбург", "city екатеренбург")
	keyboardButtons = append(keyboardButtons, button4)
	button5 := tgbotapi.NewInlineKeyboardButtonData("Казань", "city казань")
	keyboardButtons = append(keyboardButtons, button5)
	button6 := tgbotapi.NewInlineKeyboardButtonData("Самара", "city самара")
	keyboardButtons = append(keyboardButtons, button6)

	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[3:6]...)

	button7 := tgbotapi.NewInlineKeyboardButtonData("Нижний Новгород", "city нижний новгород")
	keyboardButtons = append(keyboardButtons, button7)
	button8 := tgbotapi.NewInlineKeyboardButtonData("Ростов", "city ростов")
	keyboardButtons = append(keyboardButtons, button8)
	button9 := tgbotapi.NewInlineKeyboardButtonData("Уфа", "city уфа")
	keyboardButtons = append(keyboardButtons, button9)

	row3 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[6:9]...)

	button10 := tgbotapi.NewInlineKeyboardButtonData("Железногорск(Курская обл.)", "city железногорск")
	keyboardButtons = append(keyboardButtons, button10)

	row4 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[9:]...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4)

	return &keyboard
}

func createInlineKeyboardWeather() *tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Погода на сегодня", "Погода 1")
	keyboardButtons = append(keyboardButtons, button1)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	button2 := tgbotapi.NewInlineKeyboardButtonData("Погода на завтра", "Погода 2")
	keyboardButtons = append(keyboardButtons, button2)
	button3 := tgbotapi.NewInlineKeyboardButtonData("Погода на 5 дней", "Погода 5")
	keyboardButtons = append(keyboardButtons, button3)
	button4 := tgbotapi.NewInlineKeyboardButtonData("Погода на 10 дней", "Погода 10")
	keyboardButtons = append(keyboardButtons, button4)

	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[2:]...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	return &keyboard
}

func createInlineKeyboardCalculator() *tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Сложение", "calc +")
	keyboardButtons = append(keyboardButtons, button1)
	button2 := tgbotapi.NewInlineKeyboardButtonData("Вычитание", "calc -")
	keyboardButtons = append(keyboardButtons, button2)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	button3 := tgbotapi.NewInlineKeyboardButtonData("Умножение", "calc *")
	keyboardButtons = append(keyboardButtons, button3)
	button4 := tgbotapi.NewInlineKeyboardButtonData("Деление", "calc /")
	keyboardButtons = append(keyboardButtons, button4)

	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[2:]...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	return &keyboard
}

func createInlineKeyboardOpenAPI() *tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Узнать курс BTC/USD", "BTC/USD")
	keyboardButtons = append(keyboardButtons, button1)
	button2 := tgbotapi.NewInlineKeyboardButtonData("Калорийность фруктов", "Калорийность фруктов")
	keyboardButtons = append(keyboardButtons, button2)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1)

	return &keyboard
}

func createInlineKeyboardFruits() *tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Яблоко", "apiFruit apple яблок")
	keyboardButtons = append(keyboardButtons, button1)
	button2 := tgbotapi.NewInlineKeyboardButtonData("Банан", "apiFruit Banana бананов")
	keyboardButtons = append(keyboardButtons, button2)
	button3 := tgbotapi.NewInlineKeyboardButtonData("Апельсин", "apiFruit Orange апельсинов")
	keyboardButtons = append(keyboardButtons, button3)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	button4 := tgbotapi.NewInlineKeyboardButtonData("Дыня", "apiFruit Melon дыни")
	keyboardButtons = append(keyboardButtons, button4)
	button5 := tgbotapi.NewInlineKeyboardButtonData("Лимон", "apiFruit Lemon лимона")
	keyboardButtons = append(keyboardButtons, button5)
	button6 := tgbotapi.NewInlineKeyboardButtonData("Лайм", "apiFruit Lime лайма")
	keyboardButtons = append(keyboardButtons, button6)

	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[3:6]...)

	button7 := tgbotapi.NewInlineKeyboardButtonData("Гранат", "apiFruit Pomegranate граната")
	keyboardButtons = append(keyboardButtons, button7)
	button8 := tgbotapi.NewInlineKeyboardButtonData("Виноград", "apiFruit Grape винограда")
	keyboardButtons = append(keyboardButtons, button8)
	button9 := tgbotapi.NewInlineKeyboardButtonData("Авокадо", "apiFruit Avocado авокадо")
	keyboardButtons = append(keyboardButtons, button9)

	row3 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[6:]...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3)

	return &keyboard
}
