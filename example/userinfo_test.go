package main

//curl -X GET -H "X-API-KEY: t2X........AlEF6"

import (
	"encoding/json"
	"fmt"
	"testing"
)

//该函数用来发文字消息，发群消息。
//如果apikey绑定了debox用户则以debox用户的名义发送消息，具体某个用户
//如果没绑定debox账户，则以机器人的名义发送消息，chatbot

//toUserId 表示@谁，可以为空
//groupId 发到哪个群里，required
//message 发送的消息，required
//send_msg  只能发文字，可以随便填写，required

func TestUserInfot(t *testing.T) {

	xApiKey := "t2X..."
	//url
	url := "https://open.debox.pro/openapi/authorize/userinfo?user_id=lubhe7bp"
	var header = map[string]string{
		"X-API-KEY": xApiKey,
	}
	var resp = make(map[string]interface{})
	err := HttpGet2Obj(url, header, &resp)
	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}
	strRes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	fmt.Println("UserInfot_Test success." + string(strRes))

}
