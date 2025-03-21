# debox-chat-go-sdk

`debox-chat-go-sdk` 是 DeBox 聊天服务的官方 Golang SDK。

[English version](./README_EN.md) 

## 介绍

`debox-chat-go-sdk` 封装了 DeBox 社交平台的消息发送和接收功能，助力开发者高效构建 DeBox 指令机器人，共享 DeBox 生态流量红利，实现技术价值与商业价值的双重提升。

## 主要功能

- ✅ **发送消息**：支持文本、Markdown、HTML、富文本、按钮菜单等多种消息格式。
- ✅ **接收消息**：实时监听用户消息和群聊信息，并做出自动化响应。
- ✅ **处理命令**：通过自定义指令与用户进行灵活互动。
- ✅ **Webhook 扩展**：可通过 Webhook 事件，实现更复杂的功能集成。

## 文档

完整文档请参阅 [DeBox Bot Go SDK 文档](https://docs.debox.pro/GO-SDK/)。

## 试用 SDK

1. 安装 SDK：
   
    ```sh
    go get github.com/debox-pro/debox-chat-go-sdk
    ```
2. 获取 API_KEY：
   
   您可以从 [DeBox 开放平台](https://developer.debox.pro/) 获取机器人的 API_KEY。详细信息请参阅 [DeBox 命令机器人开发指南](https://docs.debox.pro/APIs/BotGuide/)。
3. 发送您的第一条消息！
    ```go
    package main

    import "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"

    func main() {
        // 使用您的 API 密钥初始化机器人
        bot, _ := boxbotapi.NewBotAPI("YOUR_BOT_API_KEY")
        
        // 创建新消息
        msg := boxbotapi.NewMessage("DEBOX_USER_ID", "private", "Hello, DeBox!")
        
        // 发送消息
        bot.Send(msg)
    }
    ``` 
4. **更多示例，请参阅 [./example/](./example/) 目录。**

## SDK 结构

该库有三个需要理解的部分：

### Configs

Config 是与单个请求相关的字段集合。例如，如果想使用 `sendMessage` ，则可以使用 `MessageConfig` 结构来配置请求。DeBox 端点和 Config 之间存在一对一的关系。它们通常具有 `send` 前缀并以 `Config` 后缀结尾的命名模式。

### Helpers

Helper 是构建常见配置的更简单方法。如：您无需创建 `MessageConfig` 结构并设置 `ChatID` 和 `Text`，您可以使用 `NewMessage` Helper 方法，它接受请求所需的两个参数。然后，您可以在创建 `MessageConfig` 后更新它的字段。Helper 通常与方法名称相同，只是将 `send` 替换为 `New`。

### Methods

Method 用于在构建 Config 后的发送操作。通常，`Request` 是需要调用的最低级别 Method。它接受一个 `Chattable` 参数，并知道如何在需要时上传文件。它返回一个 `APIResponse`，这是 Bot API 最通用的返回类型。此方法用于没有更具体返回类型的任何端点。几乎每个其他方法都返回一个 `Message`，可以使用 `Send` 获取。

同时，存在更低级别的方法，如 `MakeRequest`，它仅需要一个端点和参数，而不是接受 Config。这些主要在内部使用。如果您需要使用它们，请提交一个 issue。

## 获取帮助

有关 DeBox API 的一般问题（不特定于 Go SDK），请查看 [DeBox 开放平台支持群](https://m.debox.pro/group?id=cc0onr82)。

有关 Go SDK 的特定问题，请创建一个新 issue，并添加 `question` 标签。

## 贡献

有关贡献 SDK 的更多信息，请参阅 [CONTRIBUTING.md](./CONTRIBUTING.md)。

## 许可证

本仓库内容的使用遵循 [LICENSE](./LICENSE) 中的许可条款。
