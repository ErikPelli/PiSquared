package PiSquared

import (
	cmap "github.com/orcaman/concurrent-map"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
)

type Bot struct {
	mem cmap.ConcurrentMap // mem is the association between user chat id (key) and its status (value).
	*tb.Bot
}

func NewBot(token string) (Bot, error) {
	bot := Bot{}
	var err error

	bot.Bot, err = tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	return bot, err
}

func (bot *Bot) InitHandlers() {
	bot.mem = cmap.New()

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
		chatId := strconv.FormatInt(m.Sender.ID, 10)
		bot.mem.Set(chatId, user{s: startMessage})
	})

	bot.Handle("/next", func(m *tb.Message) {
		chatId := strconv.FormatInt(m.Sender.ID, 10)
		u, ok := bot.mem.Get(chatId)
		if !ok {
			return
		}
		userData := u.(user)
		if userData.s != subjectSelected && userData.s != waitingResponseFromUser {
			return
		}

		question, answer := getQuestion(userData.sub)
		bot.Send(m.Sender, "Your question: "+question+"\n\nSend your answer or skip using /next.", tb.ReplyMarkup{ReplyKeyboardRemove: true})

		bot.mem.Set(chatId, user{
			s:                waitingResponseFromUser,
			sub:              userData.sub,
			lastQuizQuestion: question,
			rightAnswer:      answer,
		})
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
	})
}

func (bot *Bot) handleSubject(u *tb.User, s subject) {
	bot.Send(u, "✅ The \""+s.String()+"\" school subject has been set.")
	chatId := strconv.FormatInt(u.ID, 10)
	bot.mem.Set(chatId, user{
		s:   subjectSelected,
		sub: s,
	})
	bot.Send(u, "Get a new question using /next.")
}
