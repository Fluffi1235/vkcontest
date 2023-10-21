package tg

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"log"
)

type TG struct {
	Chan tgbotapi.UpdatesChannel
	bot  *tgbotapi.BotAPI
}

func NewTG(token string) (sources.Source, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.WithMessage(err, "Error connection tg")
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "[tg NewTG]")
	}

	tg := &TG{
		Chan: updates,
		bot:  bot,
	}

	return tg, nil
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

func (tg *TG) Send(msg string, clientID int64) error {
	tgMsg := tgbotapi.NewMessage(clientID, msg)

	_, err := tg.bot.Send(tgMsg)
	if err != nil {
		return errors.WithMessagef(err, "[tg]Error send message\n")
	}
	return nil
}

func (tg *TG) SendButton(msg string, clientID int64) error {
	tgMsg := tgbotapi.NewMessage(clientID, msg)
	tgMsg.ReplyMarkup = createInlineKeyboardInfo()

	_, err := tg.bot.Send(tgMsg)
	if err != nil {
		return errors.WithMessagef(err, "[tg]Error send keyboard\n")
	}
	return nil
}

func (tg *TG) EditMessage(infoMsg *model.EditMessage) error {
	editMsg := tgbotapi.NewEditMessageText(infoMsg.ChatId, infoMsg.ButtonMessageId, infoMsg.Answer)
	_, err := tg.bot.Send(editMsg)
	if err != nil {
		return errors.WithMessagef(err, "[tg]Error edit message\n")
	}
	return nil
}

func (tg *TG) EditMessageWithButtons(answerInfo *model.EditMessageWithButtons) error {
	editMsg := tgbotapi.NewEditMessageText(answerInfo.ChatId, answerInfo.ButtonMessageId, answerInfo.Answer)
	switch answerInfo.ButtonData {
	case "мои данные":
		editMsg.ReplyMarkup = createInlineKeyboardDataUser()
	case "прогноз погоды":
		editMsg.ReplyMarkup = createInlineKeyboardWeather()
	case "калькулятор":
		editMsg.ReplyMarkup = createInlineKeyboardCalculator()
	case "open api":
		editMsg.ReplyMarkup = createInlineKeyboardOpenAPI()
	case "калорийность фруктов":
		editMsg.ReplyMarkup = createInlineKeyboardFruits()
	case "изменить город":
		editMsg.ReplyMarkup = createInlineKeyboardCity()
	}
	_, err := tg.bot.Send(editMsg)
	if err != nil {
		return errors.WithMessage(err, "[tg]Error change keyboard\n")
	}
	return nil
}
