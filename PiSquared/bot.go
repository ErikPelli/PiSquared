package PiSquared

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

type Bot struct {
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
	bot.Handle(tb.OnText, func(m *tb.Message) {
		bot.Send(m.Sender, "Welcome "+m.Chat.FirstName+"!")
	})
}
