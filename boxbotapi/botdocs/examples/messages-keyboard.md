# Inline Keyboard

This bot waits for you to send it the message "/menu" before sending you an
inline keyboard containing a URL and some numbers. When a number is clicked, it
sends you a message with your selected number.

```go
package main

import (
	"encoding/json"
	"log"
	"testing"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

var (
	TestToken = "oPM1uUmE6mIitDC8" //replace with your token
	ChatID    = "l3ixp32y"
	// TestToken       = "pPpHtOTtXsE6i5u6" //replace with your token
	// ChatID          = "ymor0jin"
	ChatType        = "group" //private,group
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

type testLogger struct {
	t *testing.T
}

func (t testLogger) Println(v ...interface{}) {
	t.t.Log(v...)
}

func (t testLogger) Printf(format string, v ...interface{}) {
	t.t.Logf(format, v...)
}

func getBot(t *testing.T) (*boxbotapi.BotAPI, error) {
	bot, err := boxbotapi.NewBotAPI(TestToken)
	bot.Debug = true

	logger := testLogger{t}
	boxbotapi.SetLogger(logger)

	if err != nil {
		t.Error(err)
	}

	return bot, err
}

func TestNewBotAPI_notoken(t *testing.T) {
	_, err := boxbotapi.NewBotAPI("")

	if err == nil {
		t.Error(err)
	}
}

func TestGetUpdates(t *testing.T) {
	bot, _ := getBot(t)

	u := boxbotapi.NewUpdate(0)

	_, err := bot.GetUpdates(u)

	if err != nil {
		t.Error(err)
	}
}

func TestSendMarkdownMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := boxbotapi.NewMessage(ChatID, ChatType, "#title,\nA test message from the test library in debox-bot-api")
	msg.ParseMode = boxbotapi.ModeMarkdownV2
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}
func TestSendHTMLMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := boxbotapi.NewMessage(ChatID, ChatType, "A test <b>html</b> <font color=\"red\">message</font><br/><a href=\"https://debox.pro\">debox</a>")
	msg.ParseMode = boxbotapi.ModeHTML
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}
func TestSendRichText(t *testing.T) {
	var imageOne = "https://data.debox.pro/dao/newpic/one.png"
	var imageTwo = "https://data.debox.pro/dao/newpic/two.png"
	var href = "https://app.debox.pro/"
	var uiImgHead = boxbotapi.UITagImg{
		Uitag:    "img",
		Src:      imageOne,
		Position: "head",
		Href:     href,
		Height:   "200",
	}
	jsonUIImgHead, _ := json.Marshal(uiImgHead)

	var uiImgFoot = boxbotapi.UITagImg{
		Uitag:    "img",
		Src:      imageTwo,
		Position: "foot",
		Href:     href,
		Height:   "300",
	}
	uiImgFootJson, _ := json.Marshal(uiImgFoot)

	var uiA = boxbotapi.UITagA{
		Uitag: "a",
		Text:  "DeBox",
		Href:  href,
	}
	uiAJson, _ := json.Marshal(uiA)
	content := "richtext https://debox.pro " + string(jsonUIImgHead) + string(uiImgFootJson) + string(uiAJson)
	//发送
	bot, _ := getBot(t)

	msg := boxbotapi.NewMessage(ChatID, ChatType, content)
	// msg.ParseMode = boxbotapi.ModeRichText
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}

func TestGetAndSend_Messages(t *testing.T) {
	// bot, err := boxbotapi.NewBotAPI(os.Getenv("DEBOX_APITOKEN"))
	bot, err := boxbotapi.NewBotAPI(TestToken)
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
			case "help":
				msg.Text = "I understand /sayhi and /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
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
