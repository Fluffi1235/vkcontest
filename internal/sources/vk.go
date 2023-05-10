package sources

import (
	"context"
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/SevereCloud/vksdk/v2/object"
	"log"
	"strconv"
)

type Paylod struct {
	Button []byte
}

type VK struct {
	VkChan chan *longpoll.LongPoll
	Bot    *api.VK
}

func NewVK(token string) Source {
	vk := api.NewVK(token)
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	updatesVK, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	Vk := &VK{}

	chn := make(chan *longpoll.LongPoll, 1)

	chn <- updatesVK

	Vk = &VK{
		VkChan: chn,
		Bot:    vk,
	}
	return Vk
}

func (vk *VK) Read(ctx context.Context, msgChan chan<- *model.Message) {
	for update := range vk.VkChan {
		print()

		select {
		case <-ctx.Done():
			return
		default:
		}
		if update != nil {
			update.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
				log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)
				userID := obj.Message.FromID
				userIDs := []string{strconv.Itoa(userID)}
				vkUsers := api.Params{}
				vkUsers["user_ids"] = userIDs
				userInfo, err := vk.Bot.UsersGet(vkUsers)
				if err != nil {
					log.Printf("Error getting user info: %v", err)
				}
				if obj.Message.Text != "" {
					if len(userInfo) > 0 {
						msg := &model.Message{
							Source:    vk.GetSource(),
							Platform:  "vk",
							Text:      obj.Message.Text,
							ChatID:    int64(obj.Message.PeerID),
							FirstName: userInfo[0].FirstName,
							LastName:  userInfo[0].LastName,
						}

						_ = msg

						msgChan <- msg
					}
				}
			})
			update.MessageEvent(func(_ context.Context, obj events.MessageEventObject) {
				if obj.Payload != nil {
					buttonbit, err := obj.Payload.MarshalJSON()
					if err != nil {
						log.Println(nil)
					}

					fmt.Println(string(buttonbit[11 : len(buttonbit)-2]))

					msg := &model.Message{
						Source:          vk.GetSource(),
						Platform:        "vk",
						ChatID:          int64(obj.UserID),
						ButtonDate:      string(buttonbit[11 : len(buttonbit)-2]),
						ButtonMessageID: obj.ConversationMessageID,
					}

					_ = msg

					msgChan <- msg
				}
			})
		}
		log.Println("Start Long Poll")
		if err := update.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func (vk *VK) GetSource() model.SourceType {
	return model.Vk
}

func (vk *VK) Send(msg string, clientID int64) {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))
	_, err := vk.Bot.MessagesSend(b.Params)
	if err != nil {
		log.Println(err)
	}
}

func (vk *VK) SendButton(msg string, clientID int64) {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))
	b.Keyboard(CreateKeyboardvk())
	_, err := vk.Bot.MessagesSend(b.Params)
	if err != nil {
		log.Println(err)
	}
}

func (vk *VK) EditMessageWithButtons(msg string, clientID int64, button string, msgId int) {
	b := params.NewMessagesSendBuilder()
	switch button {
	case "мои данные":
		b.Keyboard(createInlineKeyboardDatavk())
	case "прогноз погоды":
		b.Keyboard(createInlineKeyboardWeathervk())
	case "калькулятор":
		b.Keyboard(createInlineKeyboardCalculatorvk())
	case "open api":
		b.Keyboard(createInlineKeyboardOpenAPIvk())
	case "калорийность фруктов":
		b.Keyboard(createInlineKeyboardFruitsvk())
	case "изменить город":
		b.Keyboard(createInlineKeyboardCityvk())
	}
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))
	_, err := vk.Bot.MessagesSend(b.Params)
	if err != nil {
		log.Println(err)
	}
}

func (vk *VK) EditMessage(msg string, clientID int64, msgId int) {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))
	_, err := vk.Bot.MessagesSend(b.Params)
	if err != nil {
		log.Println(err)
	}
}

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
					Payload: "{\"button\": \"10\"}",
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
		{
			object.MessagesKeyboardButton{
				Action: object.MessagesKeyboardButtonAction{
					Type:    "callback",
					Label:   "Уфа",
					Payload: "{\"button\": \"city Уфа\"}",
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
