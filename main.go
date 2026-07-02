package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func getWeatherForecast(city string) string {
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&units=metric&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "❌ Connection error"
	}
	defer resp.Body.Close()

	var forecastData struct {
		List []struct {
			Dt   int64 `json:"dt"`
			Main struct {
				Temp     float64 `json:"temp"`
				Humidity int     `json:"humidity"`
			} `json:"main"`
			Weather []struct {
				Description string `json:"description"`
			} `json:"weather"`
			Wind struct {
				Speed float64 `json:"speed"`
			} `json:"wind"`
		} `json:"list"`
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&forecastData)
	if err != nil {
		return "❌ City not found"
	}

	if len(forecastData.List) == 0 {
		return "❌ City not found"
	}

	result := fmt.Sprintf("🌤️ 24-hour forecast for %s:\n\n", city)

	for i := 0; i < 5 && i < len(forecastData.List); i++ {
		forecast := forecastData.List[i]
		hour := (forecast.Dt % 86400) / 3600

		result += fmt.Sprintf(
			"⏰ %02d:00 - Temp: %.1f°C | %s | Humidity: %d%%\n",
			hour,
			forecast.Main.Temp,
			forecast.Weather[0].Description,
			forecast.Main.Humidity,
		)
	}

	return result
}

func main() {
	godotenv.Load()
	token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bot is started as @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for update := range bot.GetUpdatesChan(u) {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if update.Message.Command() == "start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"👋 Hi! Send me a city name to get weather\n\nExample: Kyiv, Lviv, Kharkiv")
				bot.Send(msg)
				continue
			}
		}

		region := update.Message.Text
		log.Printf("User requested weather for: %s", region)

		weather := getWeatherForecast(region)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, weather)
		bot.Send(msg)
	}
}
