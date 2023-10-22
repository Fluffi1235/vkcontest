package parsers

import (
	"context"
	"encoding/json"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/pkg/errors"
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

func ParseFruit(msg string, config *config.Config) (*Nutrit, error) {
	fruit := Fruit{}
	url := config.FruityViceLink + msg
	client := http.DefaultClient
	client.Timeout = config.ClientTimeout
	ctx, cancel := context.WithTimeout(context.Background(), config.ContextTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.WithMessagef(err, "Error NewRequestWithContext %s", url)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithMessagef(err, "Error connecting to %s, status code error: %d %s\n", url, resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&fruit)
	if err != nil {
		return nil, errors.WithMessagef(err, "Error in data mapping %s", url)
	}
	return &fruit.Nutritions, nil
}
