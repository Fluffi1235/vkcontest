package main

import (
	"context"
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/bot"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/parse"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/Fluffi1235/vkcontest/internal/sources"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	ConnectDb string `yaml:"connectdb"`
	TgToken   string `yaml:"tgtoken"`
	VkToken   string `yaml:"vktoken"`
}

func LoadConfigFromYaml() (*Config, error) {
	var conf Config
	f, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func main() {
	config, err := LoadConfigFromYaml()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", config.ConnectDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.New(db)

	ctx := context.Background()
	var (
		mybot = bot.NewBot(map[model.SourceType]sources.Source{
			model.Telegram: sources.NewTG(config.TgToken),
			model.Vk:       sources.NewVK(config.VkToken),
		})
	)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go parse.New(repo).ParsWeather()
	go mybot.RunBot(ctx, wg, repo)

	wg.Wait()
}
