package vk

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/pkg/errors"
	"log"
	"strconv"
)

type Paylod struct {
	Button []byte
}

type VK struct {
	VkChan chan *longpoll.LongPoll
	bot    *api.VK
}

func NewVK(token string) (sources.Source, error) {
	vk := api.NewVK(token)
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Println("Error connection vk")
	}

	updatesVK, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Println(err)
	}
	Vk := &VK{}

	chn := make(chan *longpoll.LongPoll, 1)

	chn <- updatesVK

	Vk = &VK{
		VkChan: chn,
		bot:    vk,
	}

	return Vk, nil
}

func (vk *VK) Read(ctx context.Context, msgChanText chan<- *model.MessageInfoText, msgChanButton chan<- *model.MessageInfoButton) {
	for update := range vk.VkChan {
		print()

		select {
		case <-ctx.Done():
			return
		default:
		}
		if update != nil {
			update.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
				userID := obj.Message.FromID
				userIDs := []string{strconv.Itoa(userID)}
				vkUsers := api.Params{}
				vkUsers["user_ids"] = userIDs
				userInfo, err := vk.bot.UsersGet(vkUsers)
				if err != nil {
					log.Printf("[vk]Error getting message info %s: %v\n", vk.GetSource(), err)
				}
				if obj.Message.Text != "" {
					if len(userInfo) > 0 {
						msg := model.NewMessageInfoText(
							vk.GetSource(),
							"vk",
							int64(obj.Message.PeerID),
							obj.Message.Text,
							"",
							userInfo[0].FirstName,
							userInfo[0].LastName,
						)
						msgChanText <- msg
					}
				}
			})
			update.MessageEvent(func(_ context.Context, obj events.MessageEventObject) {
				if obj.Payload != nil {
					buttonBit, err := obj.Payload.MarshalJSON()
					if err != nil {
						log.Printf("[vk]Error getting buttons info %s: %v\n", vk.GetSource(), err)
					}
					msg := model.NewMessageInfoButton(
						vk.GetSource(),
						"vk",
						int64(obj.UserID),
						string(buttonBit[11:len(buttonBit)-2]),
						obj.ConversationMessageID,
					)
					msgChanButton <- msg
				}
			})
		}
		log.Println("Start Long Poll")
		if err := update.Run(); err != nil {
			log.Println(err)
		}
	}
}

func (vk *VK) GetSource() model.SourceType {
	return model.Vk
}

func (vk *VK) Send(msg string, clientID int64) error {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))
	_, err := vk.bot.MessagesSend(b.Params)
	if err != nil {
		log.Printf("Error send message %s\n", vk.GetSource())
		return errors.WithMessagef(err, "[vk]Error send message\n")
	}
	return nil
}

func (vk *VK) SendButton(msg string, clientID int64) error {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))
	b.Keyboard(CreateKeyboardvk())
	_, err := vk.bot.MessagesSend(b.Params)
	if err != nil {
		return errors.WithMessagef(err, "[vk]Error send keyboard\n")
	}
	return nil
}

func (vk *VK) EditMessage(infoMsg *model.EditMessage) error {
	b := params.NewMessagesSendBuilder()
	b.Message(infoMsg.Answer)
	b.RandomID(0)
	b.PeerID(int(infoMsg.ChatId))
	_, err := vk.bot.MessagesSend(b.Params)
	if err != nil {
		return errors.WithMessagef(err, "[vk]Error edit message\n")
	}
	return nil
}

func (vk *VK) EditMessageWithButtons(answerInfo *model.EditMessageWithButtons) error {
	b := params.NewMessagesSendBuilder()
	switch answerInfo.ButtonData {
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
	b.Message(answerInfo.Answer)
	b.RandomID(0)
	b.PeerID(int(answerInfo.ChatId))
	_, err := vk.bot.MessagesSend(b.Params)
	if err != nil {
		return errors.WithMessage(err, "[vk]Error change keyboard\n")
	}
	return nil
}
