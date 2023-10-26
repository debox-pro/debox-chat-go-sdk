package main

import (
	"fmt"
	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

func main() {

	xApiKey := ""
	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	toUserId := ""   //接收者id
	groupId := ""    //群组id
	title := ""      //消息标题
	content := ""    //消息内容
	objectName := "" //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）
	message := ""    //图文消息链接
	_, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
