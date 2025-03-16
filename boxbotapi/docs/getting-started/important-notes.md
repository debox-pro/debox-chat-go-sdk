# Important Notes

The DeBox Bot API has a few potentially unanticipated behaviors. Here are a
few of them. If any behavior was surprising to you, please feel free to open a
pull request!

## Callback Queries

- Every callback query must be answered, even if there is nothing to display to
  the user. Failure to do so will show a loading icon on the keyboard until the
  operation times out.


## GetUpdatesChan

- This method is very basic and likely unsuitable for production use. Consider
  creating your own implementation instead, as it's very simple to replicate.
- This method only allows your bot to process one update at a time. You can
  spawn goroutines to handle updates concurrently or switch to webhooks instead.
  Webhooks are suggested for high traffic bots.

## Nil Updates

- At most one of the fields in an `Update` will be set to a non-nil value. When
  evaluating updates, you must make sure you check that the field is not nil
  before trying to access any of it's fields.

## Privacy Mode

- By default, bots only get updates directly addressed to them. If you need to
  get all messages, you must disable privacy mode with Botfather. Bots already
  added to groups will need to be removed and readded for the changes to take
  effect. You can read more on the [DeBox Bot API docs][api-docs].

## User and Chat ID size

- They are `string` types
  in this library, but make sure you use correct types when saving them to a
  database or passing them to another language.
