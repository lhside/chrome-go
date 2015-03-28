# Chrome native messaging host in go

This is a package to communicate with Chrome extension using native messaging api in go.

## Install

```bash
go install github.com/lhside/chrome-go
```

## Example

Read message from standard input.

```go
msg, err := chrome.Recieve(os.Stdin)
```

Post message to standard output
```go
err := chrome.Post(msg, os.Stdout)
```

## SEE

1. [Native Messaging - Google Chrome](https://developer.chrome.com/extensions/nativeMessaging)
