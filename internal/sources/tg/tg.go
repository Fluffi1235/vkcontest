package tg

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type TG struct {
	Chan tgbotapi.UpdatesChannel
	bot  *tgbotapi.BotAPI
}

func NewTG(token string) sources.Source {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Error connection tg")
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

func (tg *TG) Read(ctx context.Context, msgChanText chan<- *model.MessageInfoText, msgChanButton chan<- *model.MessageInfoButton) {
	for update := range tg.Chan {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if update.Message != nil {
			msg := model.NewMessageInfoText(
				tg.GetSource(),
				"tg",
				update.Message.Chat.ID,
				update.Message.Text,
				update.Message.Chat.UserName,
				update.Message.Chat.FirstName,
				update.Message.Chat.LastName,
			)
			_ = msg
			msgChanText <- msg
		}
		if update.CallbackQuery != nil {
			msg := model.NewMessageInfoButton(tg.GetSource(), "tg", int64(update.CallbackQuery.From.ID), update.CallbackQuery.Data, update.CallbackQuery.Message.MessageID)
			_ = msg
			msgChanButton <- msg
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
		log.Printf("Error send message %s\n", tg.GetSource())
	}
}

func (tg *TG) SendButton(msg string, clientID int64) {
	tgMsg := tgbotapi.NewMessage(clientID, msg)
	tgMsg.ReplyMarkup = createInlineKeyboardInfo()

	_, err := tg.bot.Send(tgMsg)
	if err != nil {
		log.Printf("Error send keyboard %s\n", tg.GetSource())
	}

}

func (tg *TG) EditMessage(infoMsg *model.EditMessage) {
	editmsg := tgbotapi.NewEditMessageText(infoMsg.ChatId, infoMsg.ButtonMessageId, infoMsg.Answer)
	_, err := tg.bot.Send(editmsg)
	if err != nil {
		log.Printf("Error edit message %s\n", tg.GetSource())
	}
}

func (tg *TG) EditMessageWithButtons(answerInfo *model.EditMessageWithButtons) {
	editmsg := tgbotapi.NewEditMessageText(answerInfo.ChatId, answerInfo.ButtonMessageId, answerInfo.Answer)
	switch answerInfo.ButtonData {
	case "мои данные":
		editmsg.ReplyMarkup = createInlineKeyboardDataUser()
	case "прогноз погоды":
		editmsg.ReplyMarkup = createInlineKeyboardWeather()
	case "калькулятор":
		editmsg.ReplyMarkup = createInlineKeyboardCalculator()
	case "open api":
		editmsg.ReplyMarkup = createInlineKeyboardOpenAPI()
	case "калорийность фруктов":
		editmsg.ReplyMarkup = createInlineKeyboardFruits()
	case "изменить город":
		editmsg.ReplyMarkup = createInlineKeyboardCity()
	}
	_, err := tg.bot.Send(editmsg)
	if err != nil {
		log.Printf("Error change keyboard %s\n", tg.GetSource())
	}
}
