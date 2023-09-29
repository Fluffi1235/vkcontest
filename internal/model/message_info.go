package model

type MessageInfo struct {
	Source   SourceType
	Platform string
	ChatID   int64
}

type MessageInfoText struct {
	Mi        MessageInfo
	Text      string
	FirstName string
	LastName  string
	UserName  string
}

type MessageInfoButton struct {
	Mi              MessageInfo
	ButtonMessageID int
	ButtonData      string
}

type EditMessage struct {
	Answer          string
	ChatId          int64
	ButtonMessageId int
}

type EditMessageWithButtons struct {
	Answer          string
	ChatId          int64
	ButtonMessageId int
	ButtonData      string
}

func NewMessageInfoText(source SourceType, platform string, chatId int64, text, userName, firstName, lastName string) *MessageInfoText {
	return &MessageInfoText{
		Mi: MessageInfo{
			Source:   source,
			Platform: platform,
			ChatID:   chatId,
		},
		Text:      text,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
	}
}

func NewMessageInfoButton(source SourceType, platform string, chatId int64, buttonDate string, ButtonMessageid int) *MessageInfoButton {
	return &MessageInfoButton{
		Mi: MessageInfo{
			Source:   source,
			Platform: platform,
			ChatID:   chatId,
		},
		ButtonData:      buttonDate,
		ButtonMessageID: ButtonMessageid,
	}
}

func NewEditMessage(ans string, chatId int64, buttonMessageId int) *EditMessage {
	return &EditMessage{
		Answer:          ans,
		ChatId:          chatId,
		ButtonMessageId: buttonMessageId,
	}
}

func NewEditMessageWithButtons(ans, buttonData string, chatId int64, buttonMessageId int) *EditMessageWithButtons {
	return &EditMessageWithButtons{
		Answer:          ans,
		ButtonData:      buttonData,
		ChatId:          chatId,
		ButtonMessageId: buttonMessageId,
	}
}
