package main

import (
	"fmt"
	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

func main() {

	xApiKey := "xxxxxx"
	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	toUserId := ""
	objectName := ""
	message := ""
	_, err := client.SendRobotMsg(toUserId, message, objectName, "send_robot_msg")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
