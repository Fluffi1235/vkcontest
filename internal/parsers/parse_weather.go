package parsers

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
	"github.com/Fluffi1235/vkcontest/internal/repository"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

type Parser struct {
	repo repository.WheatherRepo
}

func New(repo repository.WheatherRepo) *Parser {
	return &Parser{
		repo: repo,
	}
}

func (p *Parser) ParsWeather(tUpdateWeeathe int) {
	citiesLinks := model.Cities()
	ticker := time.Tick(time.Duration(tUpdateWeeathe) * time.Hour)
	for range ticker {
		p.repo.ClearDb()
		for key, value := range citiesLinks {
			resp, err := http.Get(value)
			if err != nil {
				log.Printf("Error connection %s\n", value)
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				log.Printf("status code error: %d %s\n", resp.StatusCode, resp.Status)
			}
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				log.Println(err)
			}
			date := time.Now()
			var k int
			var counter int
			doc.Find(".weather-table__row").Each(func(i int, s *goquery.Selection) {
				temp := s.Find(".a11y-hidden").First().Text()
				weather := s.Find(".weather-table__body-cell_type_condition").Text()
				pressure := s.Find(".weather-table__body-cell_type_air-pressure").Text()
				humidity := s.Find(".weather-table__body-cell_type_humidity").Text()
				windspeed := s.Find(".wind-speed").Text()
				felt := s.Find(".temp__value_with-unit").Last().Text()
				timesOfDay := strings.Split(temp, ",")[0]
				temp = strings.Split(temp, ",")[1]
				counter++
				p.repo.SaveWeather(date.AddDate(0, 0, k), timesOfDay, temp, weather, pressure, humidity, windspeed, felt, key)
				if counter == 4 {
					k++
					counter = 0
				}
			})
		}
	}
}
