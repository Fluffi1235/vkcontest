package vk

import (
	"github.com/SevereCloud/vksdk/v2/object"
)

func CreateKeyboardvk() *object.MessagesKeyboard {
	buttons := [][]object.MessagesKeyboardButton{
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Мои данные",
					Payload: "{\"button\": \"Мои данные\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Прогноз Погоды",
					Payload: "{\"button\": \"Прогноз погоды\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Калькулятор",
					Payload: "{\"button\": \"Калькулятор\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "OPEN API",
					Payload: "{\"button\": \"OPEN API\"}",
				},
				Color: "primary",
			},
		},
	}

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.Buttons = buttons
	keyboard.Inline = true

	return keyboard
}

func createInlineKeyboardDatavk() *object.MessagesKeyboard {
	buttons := [][]object.MessagesKeyboardButton{
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Показать мои данные",
					Payload: "{\"button\": \"Показать мои данные\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Изменить город",
					Payload: "{\"button\": \"Изменить город\"}",
				},
				Color: "primary",
			},
		},
	}

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.Buttons = buttons
	keyboard.Inline = true
	return keyboard
}

func createInlineKeyboardCityvk() *object.MessagesKeyboard {
	buttons := [][]object.MessagesKeyboardButton{
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Санкт-Петербург",
					Payload: "{\"button\": \"city Санкт-Петербург\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Москва",
					Payload: "{\"button\": \"city Москва\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Новосибирск",
					Payload: "{\"button\": \"city Новосибирск\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Екатеренбург",
					Payload: "{\"button\": \"city Екатеренбург\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Казань",
					Payload: "{\"button\": \"city Казань\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Самара",
					Payload: "{\"button\": \"city Самара\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Нижний Новгород",
					Payload: "{\"button\": \"city Нижний Новгород\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Ростов",
					Payload: "{\"button\": \"city Ростов\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Уфа",
					Payload: "{\"button\": \"city Уфа\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Железногорск(Курская обл.)",
					Payload: "{\"button\": \"city железногорск\"}",
				},
				Color: "primary",
			},
		},
	}

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.Buttons = buttons
	keyboard.Inline = true
	return keyboard
}

func createInlineKeyboardWeathervk() *object.MessagesKeyboard {
	buttons := [][]object.MessagesKeyboardButton{
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Погода на сегодня",
					Payload: "{\"button\": \"Погода 1\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Погода на 5 дней",
					Payload: "{\"button\": \"Погода 5\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Погода на 10 дней",
					Payload: "{\"button\": \"Погода 10\"}",
				},
				Color: "primary",
			},
		},
	}

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.Buttons = buttons
	keyboard.Inline = true
	return keyboard
}

func createInlineKeyboardCalculatorvk() *object.MessagesKeyboard {
	buttons := [][]object.MessagesKeyboardButton{
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Сложение",
					Payload: "{\"button\": \"calc +\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Вычитание",
					Payload: "{\"button\": \"calc -\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Умножение",
					Payload: "{\"button\": \"calc *\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Деление",
					Payload: "{\"button\": \"calc /\"}",
				},
				Color: "primary",
			},
		},
	}

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.Buttons = buttons
	keyboard.Inline = true
	return keyboard
}

func createInlineKeyboardOpenAPIvk() *object.MessagesKeyboard {
	buttons := [][]object.MessagesKeyboardButton{
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Узнать курс BTC/USD",
					Payload: "{\"button\": \"BTC\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Калорийность фруктов",
					Payload: "{\"button\": \"Калорийность фруктов\"}",
				},
				Color: "primary",
			},
		},
	}

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.Buttons = buttons
	keyboard.Inline = true
	return keyboard
}

func createInlineKeyboardFruitsvk() *object.MessagesKeyboard {
	buttons := [][]object.MessagesKeyboardButton{
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Яблоко",
					Payload: "{\"button\": \"apiFruit apple яблок\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Банан",
					Payload: "{\"button\": \"apiFruit Banana бананов\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Дыня",
					Payload: "{\"button\": \"apiFruit Melon дыни\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Лимон",
					Payload: "{\"button\": \"apiFruit Lemon лимона\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Лайм",
					Payload: "{\"button\": \"apiFruit Lime лайма\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Гранат",
					Payload: "{\"button\": \"apiFruit Pomegranate граната\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Виноград",
					Payload: "{\"button\": \"apiFruit Grape винограда\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Авокадо",
					Payload: "{\"button\": \"apiFruit Avocado авокадо\"}",
				},
				Color: "primary",
			},
		},
	}

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.Buttons = buttons
	keyboard.Inline = true
	return keyboard
}
