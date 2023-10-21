package main

import (
	"context"
	Bot "github.com/Fluffi1235/vkcontest/internal/bot"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parsers"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	TG "github.com/Fluffi1235/vkcontest/internal/sources/tg"
	VK "github.com/Fluffi1235/vkcontest/internal/sources/vk"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func main() {
	cfg, err := config.LoadConfigFromYaml()
	if err != nil {
		log.Fatal("Error reading config")
	}

	db, err := sqlx.Connect("postgres", cfg.ConnectDb)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer db.Close()

	repo := repository.New(db)
	wg := &sync.WaitGroup{}

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		err = parsers.New(repo).ParsWeather(cfg.WeatherUpdateInfo, wg)
		if err != nil {
			log.Println(err)
		}
	}(wg)

	ctx := context.Background()

	tg, err := TG.NewTG(cfg.TgToken)
	if err != nil {
		log.Println(err)
	}

	vk, err := VK.NewVK(cfg.VkToken)
	if err != nil {
		log.Println(err)
	}

	bot := Bot.NewBot(map[model.SourceType]sources.Source{
		model.Telegram: tg,
		model.Vk:       vk,
	})

	wg.Add(1)
	go bot.RunBot(ctx, wg, repo, cfg)
	wg.Wait()
}
