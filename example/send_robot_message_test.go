package main

import (
	"encoding/json"
	"fmt"
	"testing"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
	"github.com/debox-pro/debox-chat-go-sdk/model"
)

// 给私人发送消息
// 只能发送text消息
// objectName 必须为 "RCD:Command"
func TestSendRobotMsg(t *testing.T) {

	xApiKey := "t2X..."

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)
	toUserId := "ii0k2v5n"      //接收者id
	message := "51+60=?"        //消息内容
	objectName := "RCD:Command" //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）

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

func TestSendRobotMsg_AIGC(t *testing.T) {

	xApiKey := "t2XJ..."

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)
	toUserId := "ii0k2v5n"      //接收者id
	objectName := "RCD:Graphic" //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	content := ""               //图文消息链接

	var uiA = model.UITagA{
		Uitag: "a",
		Text:  "BTC ETH",
		Href:  "https://www.baidu.com",
	}
	jsonUIA, _ := json.Marshal(uiA)
	content = string(jsonUIA)

	_, err := client.SendRobotMsg(toUserId, content, objectName, "send_robot_msg")

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
