module naverdictionary

go 1.23.6

replace naverdictionary/telegram => ./telegram

replace naverdictionary/scraper => ./scraper

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
)
