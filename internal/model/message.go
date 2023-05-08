package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Message struct {
	Source    SourceType
	Text      string
	ChatID    int64
	Username  string
	FirstName string
	LastName  string
	Button    *tgbotapi.CallbackQuery
}
