package model

type Message struct {
	Source          SourceType
	Platform        string
	Text            string
	ChatID          int64
	Username        string
	FirstName       string
	LastName        string
	ButtonDate      string
	ButtonMessageID int
}
