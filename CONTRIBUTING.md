# Contributing

If you want to contribute to the project, here's some more information about
the internal structure of the library.

# Adding Endpoints

This is mostly useful if you've managed to catch a new DeBox Bot API update
before the library can get updated. It's also a great source of information
about how the types work internally.

## Creating the Config

The first step in adding a new endpoint is to create a new Config type for it.
These belong in `configs.go`.

Let's try and add the `deleteMessage` endpoint. We can see it requires two
fields; `chat_id` and `message_id`. We can create a struct for these.

Note that we do not support `deleteMessage` for the time being. We are just using `deleteMessage` as an example, and we will support it in the future.

```go
type DeleteMessageConfig struct {
	ChatID    ???
	MessageID int
}
```

What type should `ChatID` be? DeBox allows specifying numeric chat IDs or
channel usernames. Golang doesn't have union types, and interfaces are entirely
untyped. This library solves this by adding two fields, a `ChatID`. We can now write the struct as follows.

```go
type DeleteMessageConfig struct {
	ChannelUsername string
	ChatID          int64
	MessageID       int
}
```

Note that `ChatID` is an `string`. DeBox chat IDs can be greater than 32 bits.

Okay, we now have our struct. But we can't send it yet. It doesn't implement
`Chattable` so it won't work with `Request` or `Send`.

### Making it `Chattable`

We can see that `Chattable` only requires a few methods.

```go
type Chattable interface {
	params() (Params, error)
	method() string
}
```

`params` is the fields associated with the request. `method` is the endpoint
that this Config is associated with.

Implementing the `method` is easy, so let's start with that.

```go
func (config DeleteMessageConfig) method() string {
	return "deleteMessage"
}
```

Now we have to add the `params`. The `Params` type is an alias for
`map[string]string`. DeBox expects only a single field for `chat_id`, so we
have to determine what data to send.

We could use an if statement to determine which field to get the value from.
However, as this is a relatively common operation, there's helper methods for
`Params`. We can use the `AddFirstValid` method to go through each possible
value and stop when it discovers a valid one. Before writing your own Config,
it's worth taking a look through `params.go` to see what other helpers exist.

Now we can take a look at what a completed `params` method looks like.

```go
func (config DeleteMessageConfig) params() (Params, error) {
	params := make(Params)

	params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername)
	params.AddNonZero("message_id", config.MessageID)

	return params, nil
}
```


### Base Configs

Certain Configs have repeated elements. For example, many of the items sent to a
chat have `ChatID` fields,`ReplyMarkup`. Instead of implementing all of this
code for each item, there's a `BaseChat` that handles it for your Config.
Simply embed it in your struct to get all of those fields.

There's only a few fields required for the `MessageConfig` struct after
embedding the `BaseChat` struct.

```go
type MessageConfig struct {
	BaseChat
	Text                  string
	ParseMode             string
}
```

It also inherits the `params` method from `BaseChat`. This allows you to call
it, then you only have to add your new fields.

```go
func (config MessageConfig) params() (Params, error) {
	params, err := config.BaseChat.params()
	if err != nil {
		return params, err
	}

	params.AddNonEmpty("text", config.Text)
	// Add your other fields

	return params, nil
}
```

## Making it Friendly

After we've got a Config type, we'll want to make it more user-friendly. We can
do this by adding a new helper to `helpers.go`. These are functions that take
in the required data for the request to succeed and populate a Config.

DeBox only requires two fields to call `deleteMessage`, so this will be fast.

```go
func NewDeleteMessage(chatID int64, messageID int) DeleteMessageConfig {
	return DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	}
}
```

Sometimes it makes sense to add more helpers if there's methods where you have
to set exactly one field. You can also add helpers that accept a `username`
string for channels if it's a common operation.

And that's it! You've added a new method.
