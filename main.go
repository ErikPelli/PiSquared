package main

import (
	"github.com/ErikPelli/PiSquared/src"
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
	if err := src.LoadQuestions(questionsValue); err != nil {
		log.Fatal("Error parsing questions file")
	}

	bot, err := src.NewBot(os.Getenv("BOT_TOKEN"), os.Getenv("SQLITE_DB"))
	if err != nil {
		log.Fatal("Error create bot using the bot token and the database")
	}

	err = src.LoadModel(os.Getenv("MODEL_FOLDER"))
	if err != nil {
		log.Fatal(err)
	}
	defer src.CloseModel()

	bot.InitHandlers()
	bot.Start()
}
