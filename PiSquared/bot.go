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

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Welcome "+m.Chat.FirstName+"!")
		chatId := strconv.FormatInt(m.Sender.ID, 10)
		bot.mem.Set(chatId, user{s: startMessage})
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
	})
}
