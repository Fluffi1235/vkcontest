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
					Payload: "{\"buttons\": \"Мои данные\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Прогноз Погоды",
					Payload: "{\"buttons\": \"Прогноз погоды\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Калькулятор",
					Payload: "{\"buttons\": \"Калькулятор\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "OPEN API",
					Payload: "{\"buttons\": \"OPEN API\"}",
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
					Payload: "{\"buttons\": \"Показать мои данные\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Изменить город",
					Payload: "{\"buttons\": \"Изменить город\"}",
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
					Payload: "{\"buttons\": \"city Санкт-Петербург\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Москва",
					Payload: "{\"buttons\": \"city Москва\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Новосибирск",
					Payload: "{\"buttons\": \"city Новосибирск\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Екатеренбург",
					Payload: "{\"buttons\": \"city Екатеренбург\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Казань",
					Payload: "{\"buttons\": \"city Казань\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Самара",
					Payload: "{\"buttons\": \"city Самара\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Нижний Новгород",
					Payload: "{\"buttons\": \"city Нижний Новгород\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Ростов",
					Payload: "{\"buttons\": \"city Ростов\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Уфа",
					Payload: "{\"buttons\": \"city Уфа\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Железногорск(Курская обл.)",
					Payload: "{\"buttons\": \"city железногорск\"}",
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
					Payload: "{\"buttons\": \"Погода 1\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Погода на 5 дней",
					Payload: "{\"buttons\": \"Погода 5\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Погода на 10 дней",
					Payload: "{\"buttons\": \"Погода 10\"}",
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
					Payload: "{\"buttons\": \"calc +\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Вычитание",
					Payload: "{\"buttons\": \"calc -\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Умножение",
					Payload: "{\"buttons\": \"calc *\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Деление",
					Payload: "{\"buttons\": \"calc /\"}",
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
					Payload: "{\"buttons\": \"BTC\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Калорийность фруктов",
					Payload: "{\"buttons\": \"Калорийность фруктов\"}",
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
					Payload: "{\"buttons\": \"apiFruit apple яблок\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Банан",
					Payload: "{\"buttons\": \"apiFruit Banana бананов\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Дыня",
					Payload: "{\"buttons\": \"apiFruit Melon дыни\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Лимон",
					Payload: "{\"buttons\": \"apiFruit Lemon лимона\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Лайм",
					Payload: "{\"buttons\": \"apiFruit Lime лайма\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Гранат",
					Payload: "{\"buttons\": \"apiFruit Pomegranate граната\"}",
				},
				Color: "primary",
			},
		},
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Виноград",
					Payload: "{\"buttons\": \"apiFruit Grape винограда\"}",
				},
				Color: "primary",
			},
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Авокадо",
					Payload: "{\"buttons\": \"apiFruit Avocado авокадо\"}",
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
