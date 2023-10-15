package main

import (
	"context"
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/bot"
	"github.com/Fluffi1235/vkcontest/internal/load_configs"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parsers"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	"github.com/Fluffi1235/vkcontest/internal/sources/tg"
	"github.com/Fluffi1235/vkcontest/internal/sources/vk"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func main() {
	config, err := load_configs.LoadConfigFromYaml()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db, err := sql.Open("postgres", config.ConnectDb)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer db.Close()

	repo := repository.New(db)
	go parsers.New(repo).ParsWeather(config.WeatherUpdateInfo)

	ctx := context.Background()

	mybot := bot.NewBot(map[model.SourceType]sources.Source{
		model.Telegram: tg.NewTG(config.TgToken),
		model.Vk:       vk.NewVK(config.VkToken),
	})

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go mybot.RunBot(ctx, wg, repo, config)
	wg.Wait()
}
