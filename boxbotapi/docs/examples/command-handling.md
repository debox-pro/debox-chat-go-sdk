# Command Handling

This is a simple example of changing behavior based on a provided command.

```go
package main

import (
	"fmt"
	"log"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

func main() {
	fmt.Println("Hello, world!")
	// bot, err := boxbotapi.NewBotAPI(os.Getenv("DEBOX_APITOKEN"))
	bot, err := boxbotapi.NewBotAPI("pPpHtOTtXsE6i5u6auo57")//change this to your own token
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Name)

	u := boxbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := boxbotapi.NewMessage(update.Message.Chat.ID, update.Message.Chat.Type, "")

		// Extract the command from the Message.
		switch update.Message.Text {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
```
