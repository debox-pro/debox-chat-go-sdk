# Getting Started

This library is designed as a simple wrapper around the DeBox Bot API.
It's encouraged to read [DeBox's docs][debox-docs] first to get an
understanding of what Bots are capable of doing. They also provide some good
approaches to solve common problems.

[debox-docs]: https://core.debox.org/bots

## Installing

```bash
go get -u github.com/debox-pro/debox-chat-go-sdk/boxbotapi
```

## A Simple Bot

To walk through the basics, let's create a simple echo bot that replies to your
messages repeating what you said. Make sure you get an API token from
[@Botfather][botfather] before continuing.

Let's start by constructing a new [BotAPI][bot-api-docs].

[botfather]: https://t.me/Botfather
[bot-api-docs]: https://pkg.go.dev/github.com/debox-pro/debox-chat-go-sdk/boxbotapi?tab=doc#BotAPI

```go
package main

import (
	"os"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

func main() {
	bot, err := boxbotapi.NewBotAPI(os.Getenv("DEBOX_APITOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true
}
```

Instead of typing the API token directly into the file, we're using
environment variables. This makes it easy to configure our Bot to use the right
account and prevents us from leaking our real token into the world. Anyone with
your token can send and receive messages from your Bot!

We've also set `bot.Debug = true` in order to get more information about the
requests being sent to DeBox. If you run the example above, you'll see
information about a request to the [`getMe`][get-me] endpoint. The library
automatically calls this to ensure your token is working as expected. It also
fills in the `Self` field in your `BotAPI` struct with information about the
Bot.

Now that we've connected to DeBox, let's start getting updates and doing
things. We can add this code in right after the line enabling debug mode.

[get-me]: https://core.debox.org/bots/api#getme

```go
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure DeBox knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := boxbotapi.NewUpdate(0)

	// Tell DeBox we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling DeBox for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from DeBox.
	for update := range updates {
		// DeBox can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		msg := boxbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. DeBox can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
```

Congradulations! You've made your very own bot!

Now that you've got some of the basics down, we can start talking about how the
library is structured and more advanced features.
