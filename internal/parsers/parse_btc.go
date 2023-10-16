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

type Btc struct {
	Bpi BPI `json:"bpi"`
}

type BPI struct {
	Usd Rate `json:"USD"`
}

type Rate struct {
	Rt string `json:"rate"`
}

func ParseBtc(config *config.Config) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", config.CoinDeskLink, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса coinDesk:", err)
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error connecting to coinDesk")
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", resp.StatusCode, resp.Status)
	}
	curse := Btc{}
	err = json.NewDecoder(resp.Body).Decode(&curse)
	if err != nil {
		log.Println("Error in data mapping ParseBtc")
		return "", err
	}

	return curse.Bpi.Usd.Rt, nil
}
