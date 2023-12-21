package main

import (
	"fmt"
	"testing"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

// 给私人发送消息
// 只能发送text消息
// objectName 必须为 "RCD:Command"
func TestSendRobotMsg(t *testing.T) {

	xApiKey := "t2XJiou2Mu6AlEF6"

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	toUserId := "uvg2p6ho"           //接收者id xul
	message := "im  message content" //消息内容
	objectName := "RCD:Command"      //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）

	_, err := client.SendRobotMsg(toUserId, message, objectName, "send_robot_msg")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
