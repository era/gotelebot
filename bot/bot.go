package bot

import (
	"github.com/era/gotelebot/router"
	tb "github.com/tucnak/telebot"
	"time"
)

type TelegramBot struct {
	router router.Router
	tb     *tb.Bot
}

// setup secrets and other telebot things
func New(token string) (*TelegramBot, error) {

	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		router: router.New(),
		tb:     bot,
	}, nil
}

func (bot TelegramBot) AddRoute(command string, handler func(m *tb.Message) string) {
	bot.router.Add(command, func(m *tb.Message, b *tb.Bot) {
		b.Send(m.Sender, handler(m))
	})
}

func (bot TelegramBot) Start() {
	bot.router.Setup(bot.tb)
	bot.tb.Start()
}
