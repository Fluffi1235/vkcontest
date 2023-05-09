package main

import (
	"context"
	"sync"
	"vkcontest/internal/bot"
	"vkcontest/internal/model"
	"vkcontest/internal/parse"
	"vkcontest/internal/sources"
)

func main() {
	ctx := context.Background()
	var (
		mybot = bot.NewBot(map[model.SourceType]sources.Source{
			model.Telegram: sources.NewTG(),
			model.Vk:       sources.NewVK(),
		})
	)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go parse.ParsWeather()
	go mybot.RunBot(ctx, wg)

	wg.Wait()
}
