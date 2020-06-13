# goTelebot

A very simple wrapper around https://github.com/tucnak/telebot. This only exists because most of my personal telegram bots are super simple and I found out I was duplicating code in every single project. If you need something more complex, you will need to modify it or use the original package directly.

## Usage

1. You need a bot token, just talk with https://core.telegram.org/bots#6-botfather to get it.
1. Just setup the bot using your token and add new routes:

```go
package main

import (
	"github.com/era/gotelebot/bot"
	tb "github.com/tucnak/telebot"
)


func main() {
	bot, _ := bot.New("1234:ABCDE") // This is your bot token, put your secrets in env variables, never into code!!

	bot.AddRoute("/hello", func(m *tb.Message)string { // The first argument is the command you want to reply, second what you want to do with that command. The string returned is the message the user will receive
		return "world"
	})
  
  bot.AddRoute("/where", func(m *tb.Message)string { // add how many routes you want
		return "here"
	})

	bot.Start()
}

```


After it you just need to run `go run main.go`.
