package parsers

import (
	"context"
	"github.com/Fluffi1235/vkcontest/internal/config"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Parser struct {
	repo repository.WeatherRepo
}

func New(repo repository.WeatherRepo) *Parser {
	return &Parser{
		repo: repo,
	}
}

func (p *Parser) ParsWeather(config *config.Config, wg *sync.WaitGroup) error {
	defer wg.Done()
	citiesLinks := model.Cities()
	ticker := time.Tick(time.Duration(config.WeatherUpdateInfo) * time.Hour)
	client := http.DefaultClient
	client.Timeout = config.ClientTimeout
	ctx, cancel := context.WithTimeout(context.Background(), config.ContextTimeout)
	defer cancel()
	for range ticker {
		err := p.repo.ClearDb()
		if err != nil {
			return err
		}
		for key, url := range citiesLinks {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return errors.WithMessagef(err, "Error NewRequestWithContext %s", url)
			}
			resp, err := client.Do(req)
			if err != nil {
				return errors.WithMessagef(err, "Error connecting to %s, status code error: %d %s\n", url, resp.StatusCode, resp.Status)
			}
			defer resp.Body.Close()
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return errors.WithMessagef(err, "Error in data mapping %s", url)
			}
			date := time.Now()
			var k int
			var counter int
			doc.Find(".weather-table__row").Each(func(i int, s *goquery.Selection) {
				temp := s.Find(".a11y-hidden").First().Text()
				weather := s.Find(".weather-table__body-cell_type_condition").Text()
				pressure := s.Find(".weather-table__body-cell_type_air-pressure").Text()
				humidity := s.Find(".weather-table__body-cell_type_humidity").Text()
				windSpeed := s.Find(".wind-speed").Text()
				felt := s.Find(".temp__value_with-unit").Last().Text()
				timesOfDay := strings.Split(temp, ",")[0]
				temp = strings.Split(temp, ",")[1]
				counter++
				w := &model.Weather{
					0,
					date.AddDate(0, 0, k),
					timesOfDay,
					temp,
					weather,
					pressure,
					humidity,
					windSpeed,
					felt,
					key}
				err = p.repo.SaveWeather(w)
				if err != nil {
					log.Println(err)
				}
				if counter == 4 {
					k++
					counter = 0
				}
			})
		}
	}
	return nil
}
