package main

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

const (
	TestToken = "pPpHtOTtXsE6i5u6"
	ChatID    = "ymor0jin"
	ChatType  = "group" //private,group
	Channel   = "@boxbotapitest"
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
	var imageOne = "https://data.debox.space/dao/newpic/one.png"
	var imageTwo = "https://data.debox.space/dao/newpic/two.png"
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
	content := "richtext" + string(jsonUIImgHead) + string(uiImgFootJson) + string(uiAJson)
	//发送
	bot, _ := getBot(t)

	msg := boxbotapi.NewMessage(ChatID, ChatType, content)
	msg.ParseMode = boxbotapi.ModeRichText
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}

func TestGetAndSend_Messages(t *testing.T) {
	bot, err := boxbotapi.NewBotAPI(TestToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Name)

	u := boxbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.Name, update.Message.Text)

		msg := boxbotapi.NewMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)

		bot.Send(msg)
	}
}
