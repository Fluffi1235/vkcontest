package main

import (
	"context"
	"sync"
	"telegram_bot/internal/bot"
	"telegram_bot/internal/model"
	"telegram_bot/internal/parse"
	"telegram_bot/internal/sources"
)

func main() {
	ctx := context.Background()
	mybot := bot.NewBot(map[model.SourceType]sources.Source{
		model.Telegram: sources.NewTG(),
		model.Vk:       sources.NewVK(),
	})

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go parse.ParsWeather()
	go mybot.RunBot(ctx, wg)

	wg.Wait()
}
