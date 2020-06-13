package router

import (
	tb "github.com/tucnak/telebot"
)

type Router struct {
	routes map[string]func(m *tb.Message, b *tb.Bot)
}

func New() Router {
	return Router{
		make(map[string]func(m *tb.Message, b *tb.Bot)),
	}
}

// Adds a new command -> function to bot routes
func (r *Router) Add(command string, handler func(m *tb.Message, b *tb.Bot)) {
	r.routes[command] = handler
}

// Setup the bot to use the routes defined previously
func (r *Router) Setup(telegramBot *tb.Bot) {
	for k, v := range r.routes {
		// Creates local copy
		command := k
		handler := v
		telegramBot.Handle(command, func(m *tb.Message) {
			handler(m, telegramBot)
		})
	}
}
