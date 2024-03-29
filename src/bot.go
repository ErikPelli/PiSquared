// Copyright (c) 2022 - 2023 Erik Pellizzon
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

package src

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type Bot struct {
	db *gorm.DB
	*tb.Bot
}

func NewBot(token, db string) (Bot, error) {
	bot := Bot{}
	var err error

	bot.Bot, err = tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return bot, err
	}

	bot.db, err = gorm.Open(sqlite.Open(db), &gorm.Config{})
	bot.db.AutoMigrate(&User{})

	return bot, err
}

func (bot *Bot) InitHandlers() {
	var (
		startMenu = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		btnCS     = startMenu.Text("🖥️ Computer Science")
		btnGeo    = startMenu.Text("🗺️️ Geography")
		btnHis    = startMenu.Text("📜️ History")
	)

	startMenu.Reply(
		startMenu.Row(btnCS),
		startMenu.Row(btnGeo),
		startMenu.Row(btnHis),
	)

	bot.Handle(&btnCS, func(m *tb.Message) {
		bot.handleSubject(m.Sender, computerScience)
	})

	bot.Handle(&btnGeo, func(m *tb.Message) {
		bot.handleSubject(m.Sender, geography)
	})

	bot.Handle(&btnHis, func(m *tb.Message) {
		bot.handleSubject(m.Sender, history)
	})

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Welcome "+m.Chat.FirstName+"!\n\n📚 Select the school subject from the menu.", startMenu)
		bot.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&User{
			ChatID: m.Sender.ID,
			S:      startMessage,
		})
	})

	getUser := func(m *tb.Message) (User, error) {
		var user User
		result := bot.db.First(&user, m.Sender.ID)
		return user, result.Error
	}

	bot.Handle("/next", func(m *tb.Message) {
		userData, err := getUser(m)
		if err != nil || userData.S == startMessage {
			return
		}

		question := getQuestion(userData.Sub)
		bot.Send(m.Sender, "Your question: "+question.Question+"\n\nSend your answer or skip using /next.", &tb.ReplyMarkup{ReplyKeyboardRemove: true})

		userData.S = waitingResponseFromUser
		userData.LastQuizQuestion = question.Question
		userData.RightAnswer = question.Answer
		bot.db.Save(&userData)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		userData, err := getUser(m)
		if err != nil || userData.S != waitingResponseFromUser {
			return
		}
		percentage := evalAnswer(userData.RightAnswer, m.Text)

		var status string
		if percentage >= 68 {
			status = "\U0001F7E2 Your answer is correct."
		} else if percentage >= 48 {
			status = "\U0001F7E1 Your answer is partially correct."
		} else {
			status = "🔴 Your answer is wrong."
		}
		status += "\nYour percentage: " + strconv.Itoa(int(percentage)) + "%"

		bot.Send(m.Sender, status)
		userData.S = subjectSelected
		bot.db.Save(&userData)
	})
}

func (bot *Bot) handleSubject(u *tb.User, s subject) {
	bot.Send(u, "✅ The \""+s.String()+"\" school subject has been set.")
	bot.db.Model(&User{ChatID: u.ID}).Updates(User{S: subjectSelected, Sub: s})
	bot.Send(u, "Get a new question using /next.")
}
