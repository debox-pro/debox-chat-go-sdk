## User Guide （中文）

[README in English](https://github.com/debox-pro/debox-chat-go-sdk/blob/master/README_EN.md)

### 基本介绍

本项目是 Debox 社交聊天服务（Chat Service）API 的 Golang 编程接口，Chat Service Rest API 的封装和实现，帮助 Golang 开发人员更快编程使用 Debox 的聊天消息服务。

详细 API 接口以及含义请参考：https://help.debox.pro/openapi_cn/a/api_method

### 安装

```
go get -u github.com/debox-pro/debox-chat-go-sdk
```

### 快速入门

**前言:** 所有的使用样例都位于[example](https://github.com/debox-pro/debox-chat-go-sdk/tree/master/example)目录下。

1. **注册回调地址**

   参考[register_url_sample.go](example/register_url.go)

   ```go
   package main

   import (
       "fmt"
       dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
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
       dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
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

3. **发送机器人消息**

   参考 [send_robot_msg_sample.go](example/send_robot_msg.go)

   ```go
   package main

   import (
       "fmt"
       dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
   )

   func main() {

       xApiKey := "xxxxxx"
       client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

       toUserId := ""
       fromUserId := ""
       objectName := ""
       message := ""
       _, err := client.SendRobotMsg(toUserId, message, objectName, "send_robot_msg")

       if err != nil {
           fmt.Println("send chat message fail:", err)
           return
       }

       fmt.Println("send chat message success.")

   }
   ```

4. **发送机器人群组消息**

   参考 [send_robot_group_msg_sample.go](example/send_robot_group_msg.go)

   ```go
   package main

   import (
       "fmt"
       dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
   )

   func main() {

        xApiKey := ""
       client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

       toUserId := ""
       groupId := ""
       title := ""
       content := ""
       objectName := ""
       message := ""
       _, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg")

       if err != nil {
   	    fmt.Println("send chat message fail:", err)
   	return
       }

       fmt.Println("send chat message success.")

   }
   ```
