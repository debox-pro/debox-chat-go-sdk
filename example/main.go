package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

const (
	// Bot token
	ImageOne = "https://data.debox.pro/dao/newpic/one.png"
	ImageTwo = "https://data.debox.pro/dao/newpic/two.png"
	Href     = "https://app.debox.pro/"
)

var (
	// Menu texts
	firstMenu  = "<b>Menu 1</b><br/>A box button message."
	secondMenu = "<b>Menu 2</b>  A box button message."

	// Button texts
	nextButton     = "Next"
	backButton     = "Back"
	tutorialButton = "Box"
	tokenUrl       = "https://deswap.pro/?from_chain_id=1&from_address=0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE&to_chain_id=1&to_address=0x32b77729cd87f1ef2bea4c650c16f89f08472c69&native=true"

	// Store bot screaming status
	bot *boxbotapi.BotAPI

	// Keyboard layout for the first menu. One button, one row
	firstMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonURL("url1", tokenUrl),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(nextButton, nextButton),
			boxbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	secondMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonURL(tutorialButton, tokenUrl),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BTC", "", tokenUrl, "61", "#ff0000"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BNB", "", tokenUrl, "27.5%", "#00ff00"),
			boxbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
	)
	//you can set font size and font color as follows
	// secondMenuMarkup.FontSize = "s"
	// secondMenuMarkup.FontColor = "#0000ff"
	// firstMenuMarkup.FontSize = "s"
	// firstMenuMarkup.FontColor = "#ff0000"

	validHTMLFormat = `
	<span style="color:red">span123</span>
	<b>bold</b>,nobold <strong>bold</strong>
	<i>italic</i>, <em>italic</em>
	<u>underline</u>, <ins>underline</ins>
	<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
	<span style="">spoiler</span>,
	<b>bold <i>italic bold <s>italic bold strikethrough <span>italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>
	<a href="http://www.example.com/">inline URL</a>
	<a href="box://user?id=123456789">inline mention of a user</a>
	<a href="https://debox.pro">debox</a>
	<code>inline fixed-width code</code>
	<pre>pre-formatted fixed-width code block</pre>
	<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>
	<blockquote>Block quotation started\nBlock quotation continued\nThe last line of the block quotation</blockquote>
	<blockquote expandable>Expandable block quotation started\nExpandable block quotation continued\nExpandable block quotation continued\nHidden by default part of the block quotation started\nExpandable block quotation continued\nThe last line of the block quotation</blockquote>
	`
	validMarkdownFormat = "*粗斜体*,\n" +
		"**粗斜体**,\n" +
		"~~strikethrough~~\n" +
		"# 一级标题。\n" +
		"[debox](https://debox.pro/)\n" +
		"## 22222222BTC\n" +
		"### 3333333BTC\n" +
		"#### 44444BTC\n" +
		"##### 55555555BTC\n" +
		"###### 6666666BTC\n" +
		"####### 7777777$BOX"
	contentNormal = "$box"

	htmlSample = `
	<body style="background-color1: #ff0000;" href="https://deswap.pro/?from_chain_id=56&from_address=0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE&to_chain_id=888888&to_address=vBOX">
	
	<div style="background-color:#00ff00;width:80%;color:#0000ff">
		<img src="https://data.debox.pro/dao/newpic/one.png" style="width:50;height:50;vertical-align:middle;border-radius: 50%;radius:-1"/>张三丰
		</div>
		<span style="color1:#00ff00"><b>快来一起参与吧</b> <a href="https://debox.pro">Go!</a></span>
		<br/>
		<img src="https://data.debox.pro/dao/newpic/one.png" style="width:100%;height:100%;"/>
		<hr/>
		<div style="width:70%">
		<img src="https://data.debox.pro/dao/newpic/one.png" style="width:20;height:20;vertical-align:middle;border-radius: 50%;radius:-1"/>
		https://debox.pro
	</div>
	
  	</body>
    `
)

func main() {
	var err error
	bot, err = boxbotapi.NewBotAPI("<YOUR_BOT_TOKEN_HERE>") //replace with your token
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
	case "/html", "/html1":
		err = sendHTMLMessage(chatId, chatType)
		break
	}

	return err
}

func handleButton(query *boxbotapi.CallbackQuery) {
	var text string

	markup := boxbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	if query.Data == nextButton {
		text = secondMenu
		markup = secondMenuMarkup
	} else if query.Data == backButton {
		text = firstMenu
		markup = firstMenuMarkup
	}

	// Replace menu text and keyboard
	msg := boxbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.Chat.Type, message.MessageID, text, markup)
	msg.ParseMode = boxbotapi.ModeHTML
	bot.Send(msg)
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

func SendMarkdownMessage(chatId, chatType string) error {
	msg := boxbotapi.NewMessage(chatId, chatType, "#title,\nA test message from the test library in debox-bot-api")
	msg.ParseMode = boxbotapi.ModeMarkdownV2
	_, err := bot.Send(msg)

	if err != nil {
		log.Printf("SendMarkdownMessage error:%v\n", err)
	}
	return err
}
func sendHTMLMessage(chatId, chatType string) error {
	msg := boxbotapi.NewMessage(chatId, chatType, "A test <b>html</b> <font color=\"red\">message</font><br/><a href=\"https://debox.pro\">debox</a>")
	msg.ParseMode = boxbotapi.ModeHTML
	_, err := bot.Send(msg)

	if err != nil {
		log.Printf("SendHTMLMessage error:%v\n", err)
	}
	return err
}
func sendRichText(chatId, chatType string) error {
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

	msg := boxbotapi.NewMessage(chatId, chatType, content)
	// msg.ParseMode = boxbotapi.ModeRichText
	_, err := bot.Send(msg)

	if err != nil {
		log.Printf("SendRichText error:%v\n", err)

	}
	return err
}

func sendHtml(chatId, chatType string) error {
	var html = htmlSample
	msg := boxbotapi.NewMessage(chatId, chatType, html)
	msg.ParseMode = boxbotapi.ModeHTML
	msg.ReplyMarkup = secondMenuMarkup
	_, err := bot.Send(msg)
	return err
}
