package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

var (
	// Button texts
	homeInfoContent = `<body href="%s"><b>Bot首页</b> <br/> 我是Debox官方演示机器人，展示Bot部分能力，供开发者参考。<br/>
					1、可以通过发送<b>botmother</b>消息直接唤醒我<br/>
					2、可以点击<a href="%s">@botmother</a>进入私聊交互。<br/>
					3、该Bot代码很少共218行，去掉注释后<font color="#0000ff">源码只有196行</font>，演示了以下功能：<br/>
					• 消息监听<br/>
					• 消息发送<br/>
					• 消息编辑<br/>
					• 按钮传参<br/>
					• 静默授权<br/>
					• 充话费业务交互等功能<br/>• 同时演示了用HTML构造原生消息的能力，极大丰富了富文本承载信息的能力。<br/>
					4、您可以<a href="https://docs.debox.pro/zh/GO-SDK">下载源码</a>，基于源码开发自己的Bot服务。<br/>
					基于SDK和此Demo，开发Bot的难度和成本很低，您可以轻松搭建自己的Bot服务。
					<br/>点击下面按钮体验吧。</body>
					`
	homeButton = "首页"
	myButton   = "我是谁"
	yourButton = "你是谁"
	// boxTokenUrl    = "https://deswap.pro/?from_chain_id=1&from_address=0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE&to_chain_id=1&to_address=0x32b77729cd87f1ef2bea4c650c16f89f08472c69&native=true"
	privateChatUrl = "https://m.debox.pro/user/chat?id="
	userHomePage   = "https://m.debox.pro/card?id="

	bot *boxbotapi.BotAPI
	// Keyboard layout for the second menu. Two buttons, one per row
	homeMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonDataWithColor("", yourButton, "", yourButton, "#21C161"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("", myButton, "", myButton, "#21C161"),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(homeButton, homeButton),
		),
	)
	//you can set font size and font color as follows
	// secondMenuMarkup.FontSize = "s"
	// secondMenuMarkup.FontColor = "#0000ff"
	cardSample = `
	<body href="%s">
		<div style="background-color1:#00ff00;width:80%%;color:#8a8a8a">
			<img src="%s" style="width:10%%;height:10%%;vertical-align:middle;border-radius: 50%%;radius:-1"/>%s
		</div>
			<b>Address: </b>%s
			<br/>
			<img src="%s" style="width:100%%;height:100%%;"/>
		<div style="width:70%%">
			<font style="font-size:12px;color:#8a8a8a">🏠</font>
			<font style="font-size:12px;color:#8a8a8a">%s</font>
		</div>
	</body>
	`
)

func main() {
	var err error
	// Set this to true to log all interactions with debox servers
	boxbotapi.Debug = false
	// bot, err = boxbotapi.NewBotAPI("<YOUR_BOT_TOKEN_HERE>") //replace with your token
	boxbotapi.SetHost("")                            //you can chage host here
	bot, err = boxbotapi.NewBotAPI("oPM1uU*****DC8") //replace with your token
	if err != nil {
		// Abort if something is wrong
		log.Panic(err)
	}

	u := boxbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	// `updates` is a golang channel which receives debox updates
	updates := bot.GetUpdatesChan(u)
	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)
	// Tell the user the bot is online
	go func() {
		for {
			//keep alive
			time.Sleep(time.Second * 30)
		}
	}()
	select {}
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

	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
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
	if len(text) > 0 {
		msg := boxbotapi.NewMessage(message.Chat.ID, message.Chat.Type, message.Text)
		msg.ParseMode = boxbotapi.ModeHTML

		if strings.ToLower(text) == "botmother" || message.Chat.Type == "private" {
			msg.Text = fmt.Sprintf(homeInfoContent, privateChatUrl+bot.Self.UserId)
			setSelected(homeMenuMarkup, 0, 100)
			msg.ReplyMarkup = homeMenuMarkup
			_, err = bot.Send(msg)
		}
	}
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}
func setSelected(keyboard boxbotapi.InlineKeyboardMarkup, row, col int) {
	for i := 0; i < len(keyboard.InlineKeyboard[row]); i++ {
		keyboard.InlineKeyboard[row][i].SubTextColor = "#21C161"
		if col == i {
			keyboard.InlineKeyboard[row][i].SubTextColor = "#ff0000"
		}
	}
}
func handleButton(query *boxbotapi.CallbackQuery) {
	var text string
	message := query.Message

	markup := boxbotapi.NewInlineKeyboardMarkup()
	switch query.Data {
	case homeButton:
		var user = bot.Self
		var homePage = userHomePage + user.UserId
		text = fmt.Sprintf(homeInfoContent, homePage, privateChatUrl+user.UserId)
		setSelected(homeMenuMarkup, 0, 100)
		markup = homeMenuMarkup
	case yourButton:
		var user = bot.Self
		var homePage = userHomePage + user.UserId
		text = fmt.Sprintf(cardSample, homePage, user.Pic, user.Name, user.Address, user.Pic, homePage)
		markup = homeMenuMarkup
		setSelected(homeMenuMarkup, 0, 0)
	case myButton:
		var user = query.From
		var homePage = userHomePage + user.UserId
		text = fmt.Sprintf(cardSample, homePage, user.Pic, user.Name, user.Address, user.Pic, homePage)
		markup = homeMenuMarkup
		setSelected(homeMenuMarkup, 0, 1)
	}

	// Replace menu text and keyboard
	msg := boxbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.Chat.Type, message.MessageID, text, markup)
	msg.ParseMode = boxbotapi.ModeHTML
	bot.Send(msg)
}
