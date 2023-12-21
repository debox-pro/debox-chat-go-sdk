package main

import (
	"fmt"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

func main2() {

	xApiKey := "8umgkn795xaqjx9j"
	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

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
