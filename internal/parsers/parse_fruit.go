package parsers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"log"
	"net/http"
	"time"
)

type Fruit struct {
	Nutritions Nutrit `json:"nutritions"`
}

type Nutrit struct {
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Sugar         float64 `json:"sugar"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
}

func ParseFruit(msg string, config *config.Config) (*Nutrit, error) {
	fruit := Fruit{}
	url := config.FruityViceLink + msg
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса coinDesk:", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error connecting to fruityvice, status code error: %d %s\n", resp.StatusCode, resp.Status)
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&fruit)
	if err != nil {
		log.Println("Error in data mapping ParseFruit")
		return nil, err
	}
	return &fruit.Nutritions, nil
}
