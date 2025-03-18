# Inline Keyboard

This bot waits for you to send it the message "open" before sending you an
inline keyboard containing a URL and some numbers. When a number is clicked, it
sends you a message with your selected number.

```go
package main

import (
	"log"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

var (
	swapUrl         = "https://deswap.pro/?from_chain_id=1&from_address=0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE&to_chain_id=1&to_address=0x32b77729cd87f1ef2bea4c650c16f89f08472c69&native=true"
	numericKeyboard = boxbotapi.NewInlineKeyboardMarkup(
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonURL("debox", "https://debox.pro"),
			boxbotapi.NewInlineKeyboardButtonData("2", "2"),
			boxbotapi.NewInlineKeyboardButtonData("3", "3"),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BOX", "", swapUrl, "15%", "#ff0000"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BTC", "", "https://debox.pro", "27.5%", "#00ff00"),
		),
	)
)

func main() {
	// bot, err := boxbotapi.NewBotAPI(os.Getenv("DEBOX_APITOKEN"))
	bot, err := boxbotapi.NewBotAPI("oPM1uUmE6mIitDC8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Name)

	u := boxbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			msg := boxbotapi.NewMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)

			// If the message was open, add a copy of our numeric keyboard.
			switch update.Message.Text {
			case "/menu":
				msg.ReplyMarkup = numericKeyboard
				msg.ParseMode = boxbotapi.ModeRichText
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
			msg := boxbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.Chat.Type, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
```
