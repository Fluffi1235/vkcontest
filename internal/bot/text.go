package bot

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/services"
	"strings"
)

const (
	start string = "/start"
	info  string = "/info"
	help  string = "/help"
	begin string = "начать"
)

func HandlingText(b *Bot, msg *model.MessageInfoText, service *services.Repository, persons map[int64]rune, repo repository.UniversalRepo) {
	message := strings.ToLower(msg.Text)
	if answer, err := IsTwoNumbers(message, msg.MessageInfo.ChatID, persons); answer != "" && err == nil {
		b.Sources[msg.MessageInfo.Source].Send(answer, msg.MessageInfo.ChatID)
		return
	}
	switch message {
	case start:
		repo.SetUser(msg)
		b.Sources[msg.MessageInfo.Source].Send(service.AnswerStart(), msg.MessageInfo.ChatID)
	case begin:
		repo.SetUser(msg)
		b.Sources[msg.MessageInfo.Source].Send(service.AnswerStart(), msg.MessageInfo.ChatID)
	case info:
		b.Sources[msg.MessageInfo.Source].SendButton("Выберите функцию", msg.MessageInfo.ChatID)
	case help:
		b.Sources[msg.MessageInfo.Source].Send("Чтобы узнать какие команды есть введите /info", msg.MessageInfo.ChatID)
	default:
		b.Sources[msg.MessageInfo.Source].Send("Что то пошло не так, попробуйте еще раз /info", msg.MessageInfo.ChatID)
	}
}
