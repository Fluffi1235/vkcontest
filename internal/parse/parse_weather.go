package parse

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
	repo repository.UniversalRepo
}

func New(repo repository.UniversalRepo) *Parser {
	return &Parser{
		repo: repo,
	}
}

func (p *Parser) ParsWeather() {
	for {
		p.repo.ClearDb()
		citilink := model.City()
		for key, value := range citilink {
			res, err := http.Get(value)
			if err != nil {
				log.Println(err)
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Println("status code error: %d %s", res.StatusCode, res.Status)
			}
			doc, err := goquery.NewDocumentFromReader(res.Body)
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
				p.repo.SaveInBd(date.AddDate(0, 0, k), timesOfDay, temp, weather, pressure, humidity, windspeed, felt, key)
				if counter == 4 {
					k++
					counter = 0
				}
			})
		}
		time.Sleep(1 * time.Hour)
	}
}
