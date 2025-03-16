package main

import boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"

var (
	// Menu texts
	firstMenu = "<b>Menu 1</b><br/>A box button message."

	// Button texts
	nextButton     = "Next"
	backButton     = "Back"
	tutorialButton = "Tutorial"
	tokenUrl       = "https://deswap.pro/?from_chain_id=-200&from_address=11111111111111111111111111111111&to_chain_id=-200&to_address=BpykKPT9DoPy2WoZspkd7MvUb9QAPtX86ojmrg48pump"
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
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BTC", "", "61", "#ff0000"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BNB", "", "27.5%", "#00ff00"),
		),
	)
	//you can set font size and font color as follows
	// secondMenuMarkup.FontSize = "s"
	// secondMenuMarkup.FontColor = "#0000ff"
	// firstMenuMarkup.FontSize = "s"
	// firstMenuMarkup.FontColor = "#ff0000"

	validHTMLSample = `
	<span style="color:red">span123</span>
	<b>bold</b>,nobold <strong>bold</strong>
	<i>italic</i>, <em>italic</em>
	<u>underline</u>, <ins>underline</ins>
	<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
	<span class="box-spoiler">spoiler</span>, <box-spoiler>spoiler</box-spoiler>
	<b>bold <i>italic bold <s>italic bold strikethrough <span class="box-spoiler">italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>
	<a href="http://www.example.com/">inline URL</a>
	<a href="box://user?id=123456789">inline mention of a user</a>
	<a href="https://debox.pro">debox</a>
	<box-emoji emoji-id="5368324170671202286"></box-emoji>
	<code>inline fixed-width code</code>
	<pre>pre-formatted fixed-width code block</pre>
	<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>
	<blockquote>Block quotation started\nBlock quotation continued\nThe last line of the block quotation</blockquote>
	<blockquote expandable>Expandable block quotation started\nExpandable block quotation continued\nExpandable block quotation continued\nHidden by default part of the block quotation started\nExpandable block quotation continued\nThe last line of the block quotation</blockquote>
	`
	validMarkdownV2Sample = "*粗斜体*,\n" +
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
)
