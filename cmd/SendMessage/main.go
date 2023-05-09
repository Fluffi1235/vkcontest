package main

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/bot"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parse"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	"sync"
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
