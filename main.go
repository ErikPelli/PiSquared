// Copyright (c) 2022 Erik Pellizzon
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
