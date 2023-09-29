package tg

import (
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
