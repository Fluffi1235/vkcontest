package parsers

import (
	"context"
	"encoding/json"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/pkg/errors"
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
		return "", errors.Wrap(err, "Error NewRequestWithContext in ParseBtc")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.WithMessagef(err, "Error connecting to coinDesk, status code: %d %s\n", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	curse := Btc{}
	err = json.NewDecoder(resp.Body).Decode(&curse)
	if err != nil {
		return "", errors.Wrap(err, "Error in data mapping ParseBtc")
	}

	return curse.Bpi.Usd.Rt, nil
}
