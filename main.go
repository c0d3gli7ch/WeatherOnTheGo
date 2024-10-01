package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				Rain float64 `json:"chance_of_rain"`
			} `json:"hour"`
			Astro struct {
				Sunrise   string `json:"sunrise"`
				Sunset    string `json:"sunset"`
				Moonphase string `json:"moon_phase"`
			} `json:"astro"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {

	q := "Moscow"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	params := url.Values{}
	params.Add("key", "<api_key>")
	params.Add("q", q)
	params.Add("aqi", "no")
	params.Add("alerts", "no")

	baseURL, err := url.Parse("https://api.weatherapi.com/v1/forecast.json?")
	if err != nil {
		panic(err)
	}
	baseURL.RawQuery = params.Encode()
	result, err := http.Get(baseURL.String())
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()

	if result.StatusCode != 200 {
		panic("WeatherAPI Down")
	}

	content, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(content, &weather)
	if err != nil {
		panic(err)
	}

	location, current, astro, hour := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Astro, weather.Forecast.ForecastDay[0].Hour
	fmt.Printf("%s, %s: %.0fC, %s\nSunrise: %s\nSunset: %s\nMoon Phase: %s\n", location.Name, location.Country, current.TempC, current.Condition.Text, astro.Sunrise, astro.Sunset, astro.Moonphase)
	for _, hour := range hour {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf("%s - %.0fC, %.0f%%, %s\n", date.Format("15:04"), hour.TempC, hour.Rain, hour.Condition.Text)
		if hour.Rain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
