package main

import (
	"encoding/json"
	"fmt"
	"testing"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
	"github.com/debox-pro/debox-chat-go-sdk/model"
)

//该函数用来发文字消息，发群消息。
//如果apikey绑定了debox用户则以debox用户的名义发送消息，具体某个用户
//如果没绑定debox账户，则报错，发送失败

// toUserId := "uvg2p"            //接收者id
// groupId := "fxi3h"             //群组id
// title := "im title"               //消息标题
// content := "im content"           //消息内容
// objectName := "RC:TxtMsg"         //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）
// message := "im SendRobotGroupMsg" //图文消息时传入图片链接，文字消息时传入文字
// href :="" 文字消息，此参数传空即可

func TestSendRobotGroupMsg(t *testing.T) {

	xApiKey := "t2XJi..." //配置齐全,正式
	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	toUserId := "uvg2p6ho"            //接收者id
	groupId := "l3ixp32y"             //test1
	title := "im title"               //消息标题
	content := "im content"           //无用
	objectName := "RC:TxtMsg"         //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）
	message := "im SendRobotGroupMsg" //消息内容

	_, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg", "")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}

//该函数用来发图片消息，发群消息。
//如果apikey绑定了debox用户则以debox用户的名义发送消息，具体某个用户
//如果没绑定debox账户，则报错，发送失败

// toUserId := "uvg2p"                                                                                //接收者id
// groupId := "fxi3h"                                                                                 //群组id
// title := "im title"                                                                                   //消息标题
// content := "im content"                                                                               //消息内容
// objectName := "RCD:Graphic"
// href :="https://debox.pro/"   图文消息，传入跳转链接
func TestSendRobotGroupImg(t *testing.T) {

	xApiKey := "t2X..."

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	var toUserId = "x1dei8zv"
	groupId := "l3ixp32y" //群组id
	title := "im title"   //消息标题

	objectName := "RCD:Graphic" //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	message := ""               //图文消息链接

	var uiA = model.UITagA{
		Uitag: "a",
		Text:  "BTC ETH",
		Href:  "https://debox.pro/",
	}
	jsonUIA, _ := json.Marshal(uiA)
	content := string(jsonUIA)
	_, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg", "")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
