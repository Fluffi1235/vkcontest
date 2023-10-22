package bot

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/services"
	"strings"
)

const (
	start string = "/start"
	begin string = "начать"
	info  string = "/info"
	help  string = "/help"
)

func (b *Bot) HandlingText(msg *model.MessageInfoText, service *services.Repository, persons map[int64]rune, repo repository.UniversalRepo) error {
	message := strings.ToLower(msg.Text)
	answer, err := IsTwoNumbers(message, msg.MessageInfo.ChatID, persons)
	if err != nil {
		return err
	}
	if answer != "" {
		err = b.Sources[msg.MessageInfo.Source].Send(answer, msg.MessageInfo.ChatID)
		if err != nil {
			return err
		}
	}
	switch message {
	case start:
		err = repo.SetUser(msg)
		if err != nil {
			return err
		}
		err = b.Sources[msg.MessageInfo.Source].Send(service.AnswerStart(), msg.MessageInfo.ChatID)
		if err != nil {
			return err
		}
	case begin:
		err = repo.SetUser(msg)
		if err != nil {
			return err
		}
		err = b.Sources[msg.MessageInfo.Source].Send(service.AnswerStart(), msg.MessageInfo.ChatID)
		if err != nil {
			return err
		}
	case info:
		err = b.Sources[msg.MessageInfo.Source].SendButton("Выберите функцию", msg.MessageInfo.ChatID)
		if err != nil {
			return err
		}
	case help:
		err = b.Sources[msg.MessageInfo.Source].Send("Чтобы узнать какие команды есть введите /info", msg.MessageInfo.ChatID)
		if err != nil {
			return err
		}
	default:
		err = b.Sources[msg.MessageInfo.Source].Send("Что то пошло не так, попробуйте еще раз /info", msg.MessageInfo.ChatID)
		if err != nil {
			return err
		}
	}
	return nil
}
