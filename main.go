package main

import (
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("Error start bot using the bot token")
	}

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Welcome "+m.Chat.FirstName+"!")
	})

	bot.Start()
}
