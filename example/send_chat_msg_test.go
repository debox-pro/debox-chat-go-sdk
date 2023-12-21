package main

import (
	"fmt"
	"testing"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

//往群里发
//toUserId 表示@谁，可以为空
//groupId 发到哪个群里，required
//message 发送的消息，required
//send_msg  只能发文字，可以随便填写，required

func TestSendChatMsg_Text(t *testing.T) {

	xApiKey := "t2XJiou2Mu6AlEF6" //正式环境key  370400917@qq.com 绑定的是 xu2
	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	toUserId := "uvg2p6ho"
	groupId := "fxi3hqo5"
	message := "测试 SendChatMsg"
	msgType := "send_msg"
	_, err := client.SendChatMsg(toUserId, groupId, message, msgType)

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
