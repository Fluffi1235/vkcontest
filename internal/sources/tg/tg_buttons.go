package tg

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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

func createInlineKeyboardDataUser() *tgbotapi.InlineKeyboardMarkup {
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
	cities := model.Cities()
	for city, _ := range cities {
		button := tgbotapi.NewInlineKeyboardButtonData(city, "city "+city)
		keyboardButtons = append(keyboardButtons, button)
	}
	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[0:3]...)
	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[3:6]...)
	row3 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[6:9]...)
	row4 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[9:]...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4)

	return &keyboard
}

func createInlineKeyboardWeather() *tgbotapi.InlineKeyboardMarkup {
	var keyboardButtons []tgbotapi.InlineKeyboardButton

	button1 := tgbotapi.NewInlineKeyboardButtonData("Погода на сегодня", "Погода 1")
	keyboardButtons = append(keyboardButtons, button1)

	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons...)

	button2 := tgbotapi.NewInlineKeyboardButtonData("Погода на 5 дней", "Погода 5")
	keyboardButtons = append(keyboardButtons, button2)
	button3 := tgbotapi.NewInlineKeyboardButtonData("Погода на 10 дней", "Погода 10")
	keyboardButtons = append(keyboardButtons, button3)

	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[1:]...)

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
	fruits := model.Fruits()
	for fruit, link := range fruits {
		button := tgbotapi.NewInlineKeyboardButtonData(fruit, link)
		keyboardButtons = append(keyboardButtons, button)
	}
	row1 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[0:3]...)
	row2 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[3:6]...)
	row3 := tgbotapi.NewInlineKeyboardRow(keyboardButtons[6:]...)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3)

	return &keyboard
}
