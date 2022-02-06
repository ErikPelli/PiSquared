package main

import (
	"github.com/ErikPelli/PiSquared/PiSquared"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := PiSquared.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal("Error create bot using the bot token")
	}

	bot.InitHandlers()
	bot.Start()
}
