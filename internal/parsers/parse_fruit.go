package parsers

import (
	"encoding/json"
	"github.com/Fluffi1235/vkcontest/internal/load_configs"
	"log"
	"net/http"
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

func ParseFruit(msg string, config *load_configs.Config) Nutrit {
	fruit := Fruit{}
	url := config.FruityViceLink + msg
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error connecting to fruityvice")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", resp.StatusCode, resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&fruit)
	if err != nil {
		log.Println("Error in data mapping ParseFruit")
	}
	return fruit.Nutritions
}
