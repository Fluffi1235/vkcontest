package parsers

import (
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

func (p *Parser) ParsWeather(tUpdateWeather int, wg *sync.WaitGroup) error {
	defer wg.Done()
	citiesLinks := model.Cities()
	ticker := time.Tick(time.Duration(tUpdateWeather) * time.Hour)
	var doc *goquery.Document
	var resp *http.Response
	for range ticker {
		err := p.repo.ClearDb()
		if err != nil {
			return errors.Wrap(err, "[ParsWeather]")
		}
		for key, value := range citiesLinks {
			resp, err = http.Get(value)
			if err != nil {
				return errors.WithMessagef(err, "Error connection %s\n in ParsWeather, status code error: %d %s\n", value, resp.StatusCode, resp.Status)
			}
			defer resp.Body.Close()
			doc, err = goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return errors.Wrap(err, "[ParsWeather]")
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
