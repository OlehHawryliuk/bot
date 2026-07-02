# Weather Telegram Bot 🌤️

A Telegram bot that provides weather forecasts for any city using OpenWeatherMap API.

## Features
- Get 24-hour weather forecast for any city
- Real-time temperature, humidity, and weather description

## Prerequisites
- Go 1.21+
- Docker (optional)
- Telegram Bot Token (get from [@BotFather](https://t.me/botfather))
- OpenWeatherMap API Key (free at [openweathermap.org](https://openweathermap.org/api))

## Setup

### 1. Clone the repository
```bash
git clone https://github.com/OlehHawryliuk/weather-bot.git
cd weather-bot
```

### 2. Create `.env` file
```bash
cp .env.example .env
```

Then add your tokens:
TELEGRAM_TOKEN=your_telegram_token_here
WEATHER_API_KEY=your_openweathermap_api_key_here
### 3. Run locally
```bash
go mod download
go run main.go
```

### 4. Run with Docker
```bash
docker-compose up --build
```

## Usage

1. Find the bot on Telegram (search by username from @BotFather)
2. Send `/start` to begin
3. Send any city name to get weather forecast:
   - `Kyiv`
   - `Lviv`
   - `Kharkiv`

## Example Output
🌤️ 24-hour forecast for Kyiv:
⏰ 07:00 - Temp: 18.5°C | clear sky | Humidity: 55%
⏰ 10:00 - Temp: 22.1°C | light clouds | Humidity: 48%
⏰ 13:00 - Temp: 24.5°C | overcast clouds | Humidity: 45%
⏰ 16:00 - Temp: 23.2°C | scattered clouds | Humidity: 52%
⏰ 19:00 - Temp: 19.8°C | broken clouds | Humidity: 68%
## Tech Stack
- **Go** - Backend language
- **Telegram Bot API** - `github.com/go-telegram-bot-api/telegram-bot-api/v5`
- **OpenWeatherMap API** - Weather data
- **Docker** - Containerization

## API Endpoints Used
- OpenWeatherMap: `api.openweathermap.org/data/2.5/forecast`

## Error Handling
- Connection errors: `❌ Connection error`
- City not found: `❌ City not found`

## Author
[Oleh Hawryliuk](https://github.com/OlehHawryliuk)
