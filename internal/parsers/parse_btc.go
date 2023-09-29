package parsers

import (
	"encoding/json"
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

func ParseBtc() string {
	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	curse := Btc{}
	err = json.NewDecoder(resp.Body).Decode(&curse)
	if err != nil {
		log.Println(err)
	}

	return curse.Bpi.Usd.Rt
}
