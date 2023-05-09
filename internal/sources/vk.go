package sources

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"log"
)

type VK struct {
	vkStruct *longpoll.LongPoll
	bot      *api.VK
	Chan     chan<- string
	ChatID   int
}

func NewVK() Source {
	token := ""
	vk := api.NewVK(token)

	updatesVK, err := longpoll.NewLongPoll(vk, 14)
	if err != nil {
		log.Println(err)
	}
	Vk := &VK{}

	Vk = &VK{
		vkStruct: updatesVK,
		bot:      vk,
	}
	return Vk
}

func (vk *VK) Read(ctx context.Context, msgChan chan<- *model.Message) {
	vk.vkStruct.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		msg := &model.Message{
			Source: vk.GetSource(),
			Text:   obj.Message.Text,
			ChatID: int64(obj.Message.PeerID),
		}
		_ = msg

		msgChan <- msg
	})
}

func (vk *VK) Send(msg string, clientID int64) {
	b := params.NewMessagesSendBuilder()
	b.Message(msg)
	b.RandomID(0)
	b.PeerID(int(clientID))

	_, err := vk.bot.MessagesSend(b.Params)
	if err != nil {
		log.Fatal(err)
	}
}

func (vk *VK) SendButton(msg string, clientID int64) {

}

func (vk *VK) GetSource() model.SourceType {
	return model.Vk
}

func (vk *VK) EditMessageWithButtons(msg string, clientID int64, button string, msgId int) {

}

func (vk *VK) EditMessage(msg string, clientID int64, msgId int) {

}
