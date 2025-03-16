package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

func main() {
	var err error
	// bot, err = boxbotapi.NewBotAPI("<YOUR_BOT_TOKEN_HERE>")
	bot, err = boxbotapi.NewBotAPIWithClient("pPpHtOTtXsE6i5u6", boxbotapi.APIEndpoint, nil)
	if err != nil {
		// Abort if something is wrong
		log.Panic(err)
	}
	// Set this to true to log all interactions with debox servers
	bot.Debug = true

	u := boxbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// `updates` is a golang channel which receives debox updates
	updates := bot.GetUpdatesChan(u)

	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)

	// Tell the user the bot is online
	log.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

}

func receiveUpdates(ctx context.Context, updates boxbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update boxbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		break
	}
}

func handleMessage(message *boxbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	// Print to console
	log.Printf("%s wrote %s", user.Name, text)

	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, message.Chat.Type, text)
	} else if len(text) > 0 {
		msg := boxbotapi.NewMessageResponse(message)
		_, err = bot.Send(msg)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

// When we get a command, we react accordingly
func handleCommand(chatId, chatType string, command string) error {
	var err error

	switch command {
	case "/menu":
		err = sendMenu(chatId, chatType)
		break

	case "/menu2":
		err = sendMenu2(chatId, chatType)
		break
	}

	return err
}

func handleButton(query *boxbotapi.CallbackQuery) {
	//暂时不支持消息编辑
}

func sendMenu(chatId, chatType string) error {
	msg := boxbotapi.NewMessage(chatId, chatType, firstMenu)
	msg.ParseMode = boxbotapi.ModeHTML
	msg.ReplyMarkup = firstMenuMarkup
	_, err := bot.Send(msg)
	return err
}

func sendMenu2(chatId, chatType string) error {
	msg := boxbotapi.NewMessage(chatId, chatType, firstMenu)
	msg.ParseMode = boxbotapi.ModeHTML
	msg.ReplyMarkup = secondMenuMarkup
	_, err := bot.Send(msg)
	return err
}
