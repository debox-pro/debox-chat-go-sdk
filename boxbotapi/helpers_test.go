package boxbotapi

import (
	"testing"
)

func TestNewInlineKeyboardButtonLoginURL(t *testing.T) {
	result := NewInlineKeyboardButtonLoginURL("text", LoginURL{
		URL:                "url",
		ForwardText:        "ForwardText",
		BotUsername:        "username",
		RequestWriteAccess: false,
	})

	if result.Text != "text" ||
		result.LoginURL.URL != "url" ||
		result.LoginURL.ForwardText != "ForwardText" ||
		result.LoginURL.BotUsername != "username" ||
		result.LoginURL.RequestWriteAccess != false {
		t.Fail()
	}
}
