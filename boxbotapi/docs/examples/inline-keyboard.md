# Inline Keyboard

This bot waits for you to send it the message "open" before sending you an
inline keyboard containing a URL and some numbers. When a number is clicked, it
sends you a message with your selected number.

```go
package main

import (
	"log"
	"os"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

var numericKeyboard = boxbotapi.NewInlineKeyboardMarkup(
	boxbotapi.NewInlineKeyboardRow(
		boxbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		boxbotapi.NewInlineKeyboardButtonData("2", "2"),
		boxbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	boxbotapi.NewInlineKeyboardRow(
		boxbotapi.NewInlineKeyboardButtonData("4", "4"),
		boxbotapi.NewInlineKeyboardButtonData("5", "5"),
		boxbotapi.NewInlineKeyboardButtonData("6", "6"),
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

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			msg := boxbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			// If the message was open, add a copy of our numeric keyboard.
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			}

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling DeBox to show the user
			// a message with the data received.
			callback := boxbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := boxbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
```
