package sources

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/model"
)

type Source interface {
	Read(ctx context.Context, msgChanText chan<- *model.MessageInfoText, msgChanButton chan<- *model.MessageInfoButton)
	Send(msg string, clientID int64) error
	GetSource() model.SourceType
	SendButton(msg string, clientID int64) error
	EditMessage(answerInfo *model.EditMessage) error
	EditMessageWithButtons(answerInfo *model.EditMessageWithButtons) error
}
