package sources

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/SevereCloud/vksdk/v2/object"
	"log"
	"strconv"
)

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

		Bot: vk,
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
				if obj.Message.Text != "" {
					userID := obj.Message.FromID
					userIDs := []string{strconv.Itoa(userID)}
					vkUsers := api.Params{}
					vkUsers["user_ids"] = userIDs
					userInfo, err := vk.Bot.UsersGet(vkUsers)
					if err != nil {
						log.Printf("Error getting user info: %v", err)
					}

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
		}
		log.Println("Start Long Poll")
		if err := update.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func (vk *VK) Send(msg string, clientID int64) {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))
	b.Keyboard(CreateKeyboard())
	_, err := vk.Bot.MessagesSend(b.Params)
	if err != nil {
		log.Fatal(err)
	}
}

func (vk *VK) GetSource() model.SourceType {
	return model.Vk
}

func CreateKeyboard() *object.MessagesKeyboard {
	keyboard := object.NewMessagesKeyboard(false)

	// Добавляем кнопку с текстом "Кнопка 1"
	keyboard.Buttons = make([][]object.MessagesKeyboardButton, 0)

	keyboard.AddCallbackButton("Кнопка 2", "payload_button_2", "positive")
	keyboard.OneTime = true

	return keyboard
}

func (vk *VK) SendButton(msg string, clientID int64) {

}

func (vk *VK) EditMessageWithButtons(msg string, clientID int64, button string, msgId int) {

}

func (vk *VK) EditMessage(msg string, clientID int64, msgId int) {

}
