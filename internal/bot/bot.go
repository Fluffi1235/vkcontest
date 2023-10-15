package bot

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/bot/Handlings"
	"github.com/Fluffi1235/vkcontest/internal/bot/Handlings/buttons"
	"github.com/Fluffi1235/vkcontest/internal/load_configs"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/services"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	"sync"
)

type Bot struct {
	Sources map[model.SourceType]sources.Source
}

func NewBot(m map[model.SourceType]sources.Source) Bot {
	return Bot{
		Sources: m,
	}
}

func (b *Bot) RunBot(ctx context.Context, wg *sync.WaitGroup, repo repository.UniversalRepo, config *load_configs.Config) {
	defer wg.Done()
	msgChanText := make(chan *model.MessageInfoText)
	msgChanButton := make(chan *model.MessageInfoButton)
	for _, source := range b.Sources {
		go source.Read(ctx, msgChanText, msgChanButton)
	}
	defer close(msgChanText)
	defer close(msgChanButton)

	go b.HandlingMessage(msgChanText, repo)
	go b.HandlingButtonMessage(msgChanButton, repo, config)

	select {
	case <-ctx.Done():
		return
	}
}

func (b *Bot) HandlingMessage(msgChanText <-chan *model.MessageInfoText, repo repository.UniversalRepo) {
	service := services.New(repo)
	persons := make(map[int64]rune, 0)
	for msg := range msgChanText {
		go Handlings.HandlingText(b, msg, service, persons, repo)
	}
}

func (b *Bot) HandlingButtonMessage(msgChanButton <-chan *model.MessageInfoButton, repo repository.UniversalRepo, config *load_configs.Config) {
	service := services.New(repo)
	persons := make(map[int64]rune, 0)
	for msg := range msgChanButton {
		go buttons.HandlingButton(b, msg, service, persons, config)
	}
}
