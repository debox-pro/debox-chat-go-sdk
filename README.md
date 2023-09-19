## User Guide （中文）

[README  in English](https://github.com/bm777bm/debox-chat-go-sdk/blob/master/README_EN.md)

### 基本介绍

本项目是Debox社交聊天服务（Chat Service）API的Golang编程接口，Chat Service Rest API的封装和实现，帮助Golang开发人员更快编程使用Debox的聊天消息服务。

详细API接口以及含义请参考：https://help.debox.pro/openapi_cn/a/api_method

### 安装
```
go get -u github.com/bm777bm/debox-chat-go-sdk
```


### 快速入门

**前言:**   所有的使用样例都位于[example](https://github.com/bm777bm/debox-chat-go-sdk/tree/master/example)目录下。


1. **注册回调地址**

   参考[register_url_sample.go](example/register_url.go)

   ```go
   package main

   import (
       "fmt"
       dbx_chat "github.com/bm777bm/debox-chat-go-sdk"
   )
   
   func main() {
   
       registerUrl := "www.xxx.pro/get_message"
	   xApiKey := "xxxxxx"
   
       client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)
   
       _, err := client.RegisterCallbakUrl(registerUrl, "POST", "register")
   
       if err != nil {
           fmt.Println("register callback url  fail:", err)
           return
       }
   
       fmt.Println("register callback url success.")
   
   }
   ```

2. **发送会话消息**

   参考 [send_chat_msg_sample.go](example/send_chat_msg.go)

   ```go 
   package main

   import (
       "fmt"
       dbx_chat "github.com/bm777bm/debox-chat-go-sdk"
   )
   
   func main() {
   
       xApiKey := "xxxxxx"
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
   ```

