package main

import (
	"fmt"
	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

func main() {

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", "xxxxx")

	toUserId := ""
	groupId := ""
	message := ""
	_, err := client.SendChatMsg(toUserId, groupId, message, "send_msg")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
