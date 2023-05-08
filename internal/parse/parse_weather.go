package parse

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"telegram_bot/internal/model"
	"telegram_bot/internal/service"
	"time"
)

func ParsWeather() {
	for {
		service.ClearDb()
		citilink := model.City()
		for key, value := range citilink {
			res, err := http.Get(value)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
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
				service.SaveInBd(date.AddDate(0, 0, k), timesOfDay, temp, weather, pressure, humidity, windspeed, felt, key)
				if counter == 4 {
					k++
					counter = 0
				}
			})
		}
		time.Sleep(1 * time.Hour)
	}
}
