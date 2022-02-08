package main

import (
	"github.com/ErikPelli/PiSquared/PiSquared"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	questionsFile, err := os.Open(os.Getenv("QUESTIONS_FILE"))
	if err != nil {
		log.Fatal("Error loading questions file")
	}
	questionsValue, _ := ioutil.ReadAll(questionsFile)
	if err := PiSquared.LoadQuestions(questionsValue); err != nil {
		log.Fatal("Error parsing questions file")
	}

	bot, err := PiSquared.NewBot(os.Getenv("BOT_TOKEN"), os.Getenv("SQLITE_DB"))
	if err != nil {
		log.Fatal("Error create bot using the bot token and the database")
	}

	bot.InitHandlers()
	bot.Start()
}
