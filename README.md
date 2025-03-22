# debox-chat-go-sdk

`debox-chat-go-sdk` is the official Golang SDK for the DeBox Chat Service.

[Chinese version](./README_CN.md) 

## Introduction

`debox-chat-go-sdk` encapsulates the message sending and receiving functionalities of the DeBox social platform, helping developers efficiently build DeBox command bots, share DeBox ecosystem traffic dividends, and achieve dual value enhancement of technology and business.

## Key Features

- ✅ ​**Send Messages**: Supports multiple message formats including text, Markdown, HTML, rich text, and button menus.  
- ✅ ​**Receive Messages**: Monitor user messages and group chat information in real-time, and respond automatically.  
- ✅ ​**Handle Commands**: Interact flexibly with users through custom commands.  
- ✅ ​**Webhook Extensions**: Receive Webhook events to enable more complex functionality integration.  

## Documentation

For full documentation, please refer to the [DeBox Bot Go SDK Documentation](https://docs.debox.pro/GO-SDK/).

## Try out the SDK

1. Install the SDK:
   
    ```sh
    go get -u github.com/debox-pro/debox-chat-go-sdk
    ```
2. Get your API_KEY:
   
   You can obtain the Bot's API_KEY from [DeBox Open Platform](https://developer.debox.pro/). Refer to [DeBox Command Bot Development Guide](https://docs.debox.pro/APIs/BotGuide/) for more details.
3. Send your first message:
    ```go
    package main

    import "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"

    func main() {
        // Initialize the bot with your API key
        bot, _ := boxbotapi.NewBotAPI("YOUR_BOT_API_KEY")
        
        // Create a new message
        msg := boxbotapi.NewMessage("DEBOX_USER_ID", "private", "Hello, DeBox!")
        
        // Send the message
        bot.Send(msg)
    }
    ``` 
4. **For more examples, refer to the [./example/](./example/) directory.**

## SDK Structure

This library is generally broken into three components you need to understand:

### Configs

Configs are collections of fields related to a single request. For example, if one wanted to use the `sendMessage` endpoint, you could use the `MessageConfig` struct to configure the request. There is a one-to-one relationship between DeBox endpoints and configs. They generally have the naming pattern of the `send` prefix and they all end with the `Config` suffix.

### Helpers

Helpers are easier ways of constructing common Configs. Instead of having to create a `MessageConfig` struct and remember to set the `ChatID` and `Text`, you can use the `NewMessage` helper method. It takes the two required parameters for the request to succeed. You can then set fields on the resulting `MessageConfig` after its creation. They are generally named the same as method names except with `send` replaced with `New`.


### Methods

Methods are used to send Configs after they are constructed. Generally, `Request` is the lowest level method you'll have to call. It accepts a `Chattable` parameter and knows how to upload files if needed. It returns an `APIResponse`, the most general return type from the Bot API. This method is called for any endpoint that doesn't have a more specific return type. Almost every other method returns a `Message`, which you can use `Send` to obtain.

There's lower level methods such as `MakeRequest` which require an endpoint and parameters instead of accepting configs. These are primarily used internally. If you find yourself having to use them, please open an issue.

## Getting Help

For general DeBox API questions (not specific to the Go SDK), check out the [DeBox OpenPlatform support group](https://m.debox.pro/group?id=cc0onr82).

For questions specific to the Go SDK, create a new issue and tag it with a `question` label.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information on contributing to the SDK.

## License

The contents of this repository are licensed under the [LICENSE](./LICENSE).
