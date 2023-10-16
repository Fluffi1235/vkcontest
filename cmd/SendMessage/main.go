package main

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/bot"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parsers"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	"github.com/Fluffi1235/vkcontest/internal/sources/tg"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func main() {
	cfg, err := config.LoadConfigFromYaml()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db, err := sqlx.Connect("postgres", cfg.ConnectDb)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer db.Close()

	repo := repository.New(db)
	go parsers.New(repo).ParsWeather(cfg.WeatherUpdateInfo)

	ctx := context.Background()

	bot := bot.NewBot(map[model.SourceType]sources.Source{
		model.Telegram: tg.NewTG(cfg.TgToken),
		//model.Vk:       vk.NewVK(cfg.VkToken),
	})

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go bot.RunBot(ctx, wg, repo, cfg)
	wg.Wait()
}
