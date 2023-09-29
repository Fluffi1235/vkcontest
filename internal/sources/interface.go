package sources

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
)

type Source interface {
	Read(ctx context.Context, msgChanText chan<- *model.MessageInfoText, msgChanButton chan<- *model.MessageInfoButton)
	Send(msg string, clientID int64)
	GetSource() model.SourceType
	SendButton(msg string, clientID int64)
	EditMessage(answerInfo *model.EditMessage)
	EditMessageWithButtons(answerInfo *model.EditMessageWithButtons)
}
