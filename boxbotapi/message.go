package boxbotapi

import (
	"encoding/json"
)

type MarkdownV2Config struct {
	ToUserId         string                `json:"to_user_id"`
	GroupId          string                `json:"group_id"`
	Title            string                `json:"title"`
	Content          string                `json:"content"`
	ObjectName       string                `json:"object_name"`
	ReplyMarkup      *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	UserActionMarkup *InlineKeyboardMarkup `json:"user_action_markup,omitempty"`
}

func (message MarkdownV2Config) params() (Params, error) {
	return Params{}, nil
}
func (message MarkdownV2Config) method() string {
	bytes, err := json.Marshal(message)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (message MarkdownV2Config) mashal() []byte {
	bytes, err := json.Marshal(message)
	if err != nil {
		return []byte{}
	}
	return bytes
}
