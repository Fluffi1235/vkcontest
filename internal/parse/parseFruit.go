package parse

import (
	"encoding/json"
	"log"
	"net/http"
)

type Fruit struct {
	Nutritions Nutrit `json:"nutritions"`
}

type Nutrit struct {
	Calories       float64 `json:"calories"`
	Fats           float64 `json:"fats"`
	Sugar          float64 `json:"sugar"`
	Carbohyddrates float64 `json:"carbohyddrates"`
	Protein        float64 `json:"protein"`
}

func ParseFruit(msg string) Nutrit {
	fruit := Fruit{}
	url := "https://www.fruityvice.com/api/fruit/" + msg
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&fruit)
	if err != nil {
		log.Fatal(err)
	}
	return fruit.Nutritions
}
