module naverdictionary/main

go 1.23

replace naverdictionary/scraper => ./scraper

replace naverdictionary/telegram => ./telegram

require naverdictionary/scraper v0.0.0-00010101000000-000000000000 // indirect

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 // indirect
	naverdictionary/telegram v0.0.0-00010101000000-000000000000 // indirect
)
