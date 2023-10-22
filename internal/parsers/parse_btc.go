package parsers

import (
	"context"
	"encoding/json"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/pkg/errors"
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

func ParseBtc(config *config.Config) (string, error) {
	url := config.CoinDeskLink
	client := http.DefaultClient
	client.Timeout = config.ClientTimeout
	ctx, cancel := context.WithTimeout(context.Background(), config.ContextTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", errors.WithMessagef(err, "Error NewRequestWithContext in ParseBtc")
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.WithMessagef(err, "Error connecting %s, status code: %d %s\n", url, resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	curse := Btc{}
	err = json.NewDecoder(resp.Body).Decode(&curse)
	if err != nil {
		return "", errors.WithMessagef(err, "Error in data mapping %s", url)
	}

	return curse.Bpi.Usd.Rt, nil
}
