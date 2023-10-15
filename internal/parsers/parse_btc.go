package parsers

import (
	"encoding/json"
	"github.com/Fluffi1235/vkcontest/internal/load_configs"
	"log"
	"net/http"
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

func ParseBtc(config *load_configs.Config) string {
	resp, err := http.Get(config.CoinDeskLink)
	if err != nil {
		log.Println("Error connecting to coinDesk")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", resp.StatusCode, resp.Status)
	}
	curse := Btc{}
	err = json.NewDecoder(resp.Body).Decode(&curse)
	if err != nil {
		log.Println("Error in data mapping ParseBtc")
	}

	return curse.Bpi.Usd.Rt
}
