package main

//curl -X GET -H "X-API-KEY: t2X........AlEF6"

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"testing"
	"time"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

func getSignature(appSecret string) (nonce, timestamp, signature string) {
	nonceInt := rand.Int()
	nonce = strconv.Itoa(nonceInt)
	timeInt64 := time.Now().Unix()
	timestamp = strconv.FormatInt(timeInt64, 10)
	h := sha1.New()
	// _, _ = io.WriteString(h, c.AppSecret+nonce+timestamp)
	_, _ = io.WriteString(h, appSecret+nonce+timestamp)
	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

//该函数用来发文字消息，发群消息。
//如果apikey绑定了debox用户则以debox用户的名义发送消息，具体某个用户
//如果没绑定debox账户，则以机器人的名义发送消息，chatbot

//toUserId 表示@谁，可以为空
//groupId 发到哪个群里，required
//message 发送的消息，required
//send_msg  只能发文字，可以随便填写，required

func TestUserInfo(t *testing.T) {

	//url
	url := "https://open.debox.pro/openapi/user/info?user_id=uvg2p6ho"

	nonce, timestamp, signature := getSignature("app_secret")
	var headers = map[string]string{
		"X-API-KEY": "<YOUR_API_KEY_HERE>",
		"nonce":     nonce,
		"timestamp": timestamp,
		"signature": signature,
	}

	var resp = make(map[string]interface{})
	err := boxbotapi.HttpGet2Obj(url, headers, &resp)
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

func TestSetHost(t *testing.T) {
	//url
	boxbotapi.SetHost("https://open.debox.pro")

}
