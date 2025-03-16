package boxbotapi

import (
	"testing"
	"time"
)

func TestUserStringWith(t *testing.T) {
	user := User{
		UserId:       "dfdfd",
		FirstName:    "Test",
		LastName:     "Test",
		Name:         "",
		LanguageCode: "en",
		IsBot:        false,
	}

	if user.String() != "Test Test" {
		t.Fail()
	}
}

func TestUserStringWithUserName(t *testing.T) {
	user := User{
		UserId:       "fdfd",
		FirstName:    "Test",
		LastName:     "Test",
		Name:         "@test",
		LanguageCode: "en",
	}

	if user.String() != "@test" {
		t.Fail()
	}
}

func TestMessageTime(t *testing.T) {
	message := Message{Date: 0}

	date := time.Unix(0, 0)
	if message.Time() != date {
		t.Fail()
	}
}

func TestChatIsPrivate(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "private"}

	if !chat.IsPrivate() {
		t.Fail()
	}
}

func TestChatIsGroup(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "group"}

	if !chat.IsGroup() {
		t.Fail()
	}
}

func TestChatIsChannel(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "channel"}

	if !chat.IsChannel() {
		t.Fail()
	}
}

func TestChatIsSuperGroup(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "supergroup"}

	if !chat.IsSuperGroup() {
		t.Fail()
	}
}

// Ensure all configs are sendable
var (
	_ Chattable = CloseConfig{}
	_ Chattable = MessageConfig{}
)

// Ensure all RequestFileData types are correct.
