# Keyboard

This bot shows a numeric keyboard when you send a "open" message and hides it
when you send "close" message.

```go
package main

import (
	"log"
	"os"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

var numericKeyboard = boxbotapi.NewReplyKeyboard(
	boxbotapi.NewKeyboardButtonRow(
		boxbotapi.NewKeyboardButton("1"),
		boxbotapi.NewKeyboardButton("2"),
		boxbotapi.NewKeyboardButton("3"),
	),
	boxbotapi.NewKeyboardButtonRow(
		boxbotapi.NewKeyboardButton("4"),
		boxbotapi.NewKeyboardButton("5"),
		boxbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	bot, err := boxbotapi.NewBotAPI(os.Getenv("DEBOX_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := boxbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := boxbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = nil
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
```
