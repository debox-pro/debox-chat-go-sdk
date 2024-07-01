package main

import (
	"fmt"
	"testing"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

//该函数用来发文字消息，发群消息。
//如果apikey绑定了debox用户则以debox用户的名义发送消息，具体某个用户
//如果没绑定debox账户，则以机器人的名义发送消息，chatbot

//toUserId 表示@谁，可以为空
//groupId 发到哪个群里，required
//message 发送的消息，required
//send_msg  只能发文字，可以随便填写，required

func TestSendChatMsg_Text(t *testing.T) {

	xApiKey := "t2XJi..." //正式环境key  370400917@qq.com 绑定的是 xu2
	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	groupId := "l3ixp32y" //test1
	message := "测试 SendChatMsg1"
	msgType := "send_msg"
	toUserId := "uvg2p6ho"
	_, err := client.SendChatMsg(toUserId, groupId, message, msgType)

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
